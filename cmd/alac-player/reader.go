package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"unsafe"

	"github.com/jfbus/mp4"
	"github.com/xlab/alac-go/alac"
	"github.com/xlab/portaudio-go/portaudio"
)

var (
	AlacErr = errors.New("ALAC header corrupt")
	AlacTag = []byte("alac")
)

type AlacReader struct {
	buf   []byte
	debuf []byte
	in    io.ReadSeeker
	a     *alac.File

	frameN uint32
	chunkN uint32

	bytesPerSample    uint
	sampleRate        uint32
	framesTotal       uint32
	unitsTotal        uint32
	maxSamplesInFrame uint
	frameDurations    uint32Rows
	packetsPerChunk   uint32Rows
	chunkOffsets      []uint32
	packetSizes       []uint32
	packetUniformSize uint32

	// OnNextFrame gets invoked on each audio frame been processed.
	// n is the next frame number starting from 0 (first).
	OnNextFrame func(n int)
}

// Skip seconds of playback from beginning of the stream.
func (a *AlacReader) Skip(seconds int) error {
	unit := a.sampleRate * uint32(seconds)
	var frameN, unitSum uint32
	for unitSum < unit {
		// skip audio frames until duration (in units) accumulates
		unitSum += a.frameDurations.SearchLast(frameN)
		frameN++
	}
	a.chunkN, a.frameN = a.chunkOf(uint32(frameN))
	if a.OnNextFrame != nil {
		a.OnNextFrame(int(a.frameN))
	}
	// a.totalSamplesInChunk = a.packetsPerChunk.SearchFirst(a.chunkN)
	if _, err := a.in.Seek(a.chunkOffset(a.chunkN), 0); err != nil {
		return errors.New("cannot read audio data from file")
	}
	return nil
}

func (a *AlacReader) SampleRate() float64 {
	return float64(a.sampleRate)
}

func (a *AlacReader) SamplesTotal() int {
	return int(a.framesTotal)
}

func (a *AlacReader) Duration() float32 {
	return float32(a.unitsTotal) / float32(a.sampleRate)
}

// sampleSize gets the size of n-th sample packet in chunk.
func (a *AlacReader) packetSize(n uint32) uint32 {
	if len(a.packetSizes) == 0 {
		return a.packetUniformSize
	}
	return a.packetSizes[n]
}

func (a *AlacReader) chunkOffset(n uint32) int64 {
	return int64(a.chunkOffsets[n])
}

// chunkOf gets the chunk number the packet n belongs to and the aligned
// packets offset number since the start of stream.
func (a *AlacReader) chunkOf(n uint32) (chunk uint32, packets uint32) {
	for packets < n {
		packets += a.packetsPerChunk.SearchFirst(chunk)
		chunk++
	}
	return chunk, packets
}

func NewAlacReader(in io.ReadSeeker, bitDepth int, channels int) *AlacReader {
	return &AlacReader{
		in: in,
		a:  alac.Create(bitDepth, channels),
		//
		bytesPerSample: uint(bitDepth / 8),
	}
}

// DecodeContainer parses QuickTime container, gets audio track metadata, a timing table,
// and the ALAC header (magic cookie).
//
// [1] https://developer.apple.com/library/mac/documentation/QuickTime/QTFF/QTFFChap1/qtff1.html#//apple_ref/doc/uid/TP40000939-CH203-BBCGDDDF
// [2] https://alac.macosforge.org/trac/browser/trunk/ALACMagicCookieDescription.txt
func (a *AlacReader) DecodeContainer() error {
	v, err := mp4.Decode(a.in)
	if err != nil {
		return err
	}
	if v.Moov == nil || len(v.Moov.Trak) == 0 {
		return errors.New("no track found")
	}
	var mdia *mp4.MdiaBox
	for i := range v.Moov.Trak {
		m := v.Moov.Trak[i].Mdia
		if m != nil && m.Hdlr != nil && m.Hdlr.HandlerType == "soun" {
			mdia = m
		}
	}
	if mdia == nil {
		return errors.New("no audio track found")
	}
	a.sampleRate = mdia.Mdhd.Timescale
	a.unitsTotal = mdia.Mdhd.Duration

	table := mdia.Minf.Stbl

	size := table.Stsd.Size()
	buf := bytes.NewBuffer(make([]byte, 0, size))
	table.Stsd.Encode(buf)
	if len(buf.Next(0x34)) < 0x34 {
		return AlacErr
	} else if v := buf.Next(4); len(v) < 4 {
		return AlacErr
	} else {
		size := int(uint(v[0])<<24 | uint(v[1])<<16 | uint(v[2])<<8 | uint(v[3]))
		if cookie := buf.Next(size - 4); len(cookie) < size-4 {
			return AlacErr
		} else {
			log.Printf("ALAC header: %02X", cookie)
			alac.SetInfo(a.a, cookie)
			a.a.Deref()
			a.maxSamplesInFrame = a.a.Setinfo82
			// audio packet buffers
			a.buf = make([]byte, a.maxSamplesInFrame*a.bytesPerSample)
			a.debuf = make([]byte, a.maxSamplesInFrame*a.bytesPerSample)
		}
	}

	for i, count := range table.Stts.SampleCount {
		a.frameDurations = append(a.frameDurations, uint32Row{
			N:     count,
			Value: table.Stts.SampleTimeDelta[i],
		})
	}
	for i, first := range table.Stsc.FirstChunk {
		// audio frames per chunk
		a.packetsPerChunk = append(a.packetsPerChunk, uint32Row{
			N:     first - 1,
			Value: table.Stsc.SamplesPerChunk[i],
		})
	}

	a.framesTotal = table.Stsz.SampleNumber
	if size := table.Stsz.SampleUniformSize; size > 0 {
		a.packetUniformSize = size
	} else {
		a.packetSizes = table.Stsz.SampleSize
	}
	a.chunkOffsets = table.Stco.ChunkOffset

	// a.totalSamplesInChunk = a.packetsPerChunk.SearchFirst(0)
	if _, err := a.in.Seek(a.chunkOffset(0), 0); err != nil {
		return errors.New("cannot read audio data from file")
	}
	fmt.Printf("Audio duration: %.3fs\n", a.Duration())
	return nil
}

func (a *AlacReader) Close() {
	alac.Free(a.a)
}

// advanceChunk switches to the next chunk if all the audio frames
// in the current chunk have been played. Count of frames per chunk may vary.
//
// DEPRECATED: chunks usually are without gaps, otherwise additional
// heavy logic involved and the QuickTime container has more tables.
// func (a *AlacReader) advanceChunk()

func (a *AlacReader) StreamCallback(_ unsafe.Pointer, output unsafe.Pointer, sampleCount uint,
	_ *portaudio.StreamCallbackTimeInfo, _ portaudio.StreamCallbackFlags, _ unsafe.Pointer) int {

	const (
		statusContinue = int(portaudio.PaContinue)
		statusComplete = int(portaudio.PaComplete)
		statusAbort    = int(portaudio.PaAbort)
	)

	if a.frameN >= a.framesTotal {
		return statusComplete
	}
	// if a.advanceChunk() != nil {
	// 	return statusComplete
	// }
	packetSize := a.packetSize(a.frameN)
	if _, err := a.in.Read(a.buf[:packetSize]); err != nil {
		log.Println("[warn]:", err)
		return statusAbort
	}

	var size int
	alac.DecodeFrame(a.a, a.buf[:packetSize], unsafe.Pointer(&a.debuf[0]), &size)
	a.debuf = a.debuf[:size]
	// sampleCount a.k.a samples in the frame, a frame usually has
	// 4096 samples per channel, so we process (int16|int16) at a time for L|R.
	out := (*(*[1 << 32]int16)(output))[:sampleCount*audioChannels]
	for i := 0; i < len(out); i++ {
		// 2 channel 16-bit stereo
		out[i] = int16(a.debuf[2*i]) | int16(a.debuf[2*i+1])<<8
	}
	a.frameN++
	if a.OnNextFrame != nil {
		a.OnNextFrame(int(a.frameN))
	}
	// a.currentFrameInChunk++
	return statusContinue
}
