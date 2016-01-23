// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 23 Jan 2016 21:34:01 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package alac

/*
#include "alac.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// File as declared in alac/alac.h:6
type File struct {
	InputBuffer               []byte
	InputBufferBitaccumulator int32
	Samplesize                int32
	Numchannels               int32
	Bytespersample            int32
	PredicterrorBufferA       []int32
	PredicterrorBufferB       []int32
	OutputsamplesBufferA      []int32
	OutputsamplesBufferB      []int32
	UncompressedBytesBufferA  []int32
	UncompressedBytesBufferB  []int32
	SetinfoMaxSamplesPerFrame uint32
	Setinfo7a                 byte
	SetinfoSampleSize         byte
	SetinfoRiceHistorymult    byte
	SetinfoRiceInitialhistory byte
	SetinfoRiceKmodifier      byte
	Setinfo7f                 byte
	Setinfo80                 uint16
	Setinfo82                 uint32
	Setinfo86                 uint32
	Setinfo8aRate             uint32
	ref9e6cbf0e               *C.alac_file
	allocs9e6cbf0e            interface{}
}
