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
import "unsafe"

// Create function as declared in alac/alac.h:8
func Create(samplesize int32, numchannels int32) *File {
	csamplesize, _ := (C.int)(samplesize), cgoAllocsUnknown
	cnumchannels, _ := (C.int)(numchannels), cgoAllocsUnknown
	__ret := C.alac_create(csamplesize, cnumchannels)
	__v := NewFileRef(__ret)
	return __v
}

// DecodeFrame function as declared in alac/alac.h:9
func DecodeFrame(alac *File, inbuffer []byte, outbuffer unsafe.Pointer, outputsize *int32) {
	calac, _ := alac.PassRef()
	cinbuffer, _ := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&inbuffer)).Data)), cgoAllocsUnknown
	coutbuffer, _ := (unsafe.Pointer)(unsafe.Pointer(outbuffer)), cgoAllocsUnknown
	coutputsize, _ := (*C.int)(unsafe.Pointer(outputsize)), cgoAllocsUnknown
	C.alac_decode_frame(calac, cinbuffer, coutbuffer, coutputsize)
}

// SetInfo function as declared in alac/alac.h:12
func SetInfo(alac *File, inputbuffer []byte) {
	calac, _ := alac.PassRef()
	cinputbuffer, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&inputbuffer)).Data)), cgoAllocsUnknown
	C.alac_set_info(calac, cinputbuffer)
}

// AllocateBuffers function as declared in alac/alac.h:13
func AllocateBuffers(alac *File) {
	calac, _ := alac.PassRef()
	C.alac_allocate_buffers(calac)
}

// Free function as declared in alac/alac.h:14
func Free(alac *File) {
	calac, _ := alac.PassRef()
	C.alac_free(calac)
}
