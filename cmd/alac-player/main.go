package main

import (
	"log"
	"os"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	"github.com/xlab/portaudio-go/portaudio"
)

const (
	framesPerPacket = 4096
	bitDepth        = 16
	audioChannels   = 2
	sampleFormat    = portaudio.PaInt16
)

var (
	app      = cli.App("alac-player", "A player implemented in Go that can read M4A ALAC files and plays using PortAudio.")
	filename = app.StringArg("FILENAME", "", "An .m4a ALAC file to play. 44100Hz 16-bit stereo is only supported.")
	skip     = app.IntOpt("s skip", 0, "Skip the first X.X seconds of playback.")
)

func main() {
	log.SetFlags(0)
	app.Spec = "[--skip=<0.0>] FILENAME"
	app.Action = appRun
	app.Run(os.Args)
}

func appRun() {
	defer closer.Close()
	if err := portaudio.Initialize(); paError(err) {
		log.Fatalln("PortAudio init error:", err)
	}
	closer.Bind(func() {
		if err := portaudio.Terminate(); paError(err) {
			log.Println("PortAudio term error:", err)
		}
	})

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalln(err)
	}
	alacReader := NewAlacReader(f, bitDepth, audioChannels)
	closer.Bind(func() {
		alacReader.Close()
		f.Close()
	})
	if err := alacReader.DecodeContainer(); err != nil {
		log.Fatalln(err)
	}

	samplesTotal := alacReader.SamplesTotal()
	progress := pb.New(samplesTotal).SetRefreshRate(time.Second)
	progress.ShowSpeed = false
	progress.ShowCounters = false
	progress.ShowTimeLeft = false
	progress.ShowFinalTime = false
	alacReader.OnNextFrame = func(n int) {
		progress.Set(n)
	}

	skipSeconds := *skip
	switch {
	case float32(skipSeconds) >= alacReader.Duration():
		os.Exit(0)
	case skipSeconds > 0:
		alacReader.Skip(skipSeconds)
	case skipSeconds < 0:
		skipSeconds = int(alacReader.Duration()) + skipSeconds
		alacReader.Skip(skipSeconds)
	}

	var stream *portaudio.Stream
	if err := portaudio.OpenDefaultStream(&stream, 0, audioChannels, sampleFormat, alacReader.SampleRate(),
		framesPerPacket, alacReader.StreamCallback, nil); paError(err) {
		log.Fatalln("PortAudio error:", err)
	}
	closer.Bind(func() {
		if err := portaudio.CloseStream(stream); paError(err) {
			log.Println("[WARN] PortAudio error:", err)
		}
	})

	if err := portaudio.StartStream(stream); paError(err) {
		log.Fatalln("PortAudio error:", err)
	}
	closer.Bind(func() {
		if err := portaudio.StopStream(stream); paError(err) {
			log.Fatalln("[WARN] PortAudio error:", err)
		}
	})

	progress.Start()
	dur := time.Duration(alacReader.Duration()) * time.Second
	<-time.Tick(dur - time.Duration(skipSeconds)*time.Second)
}
