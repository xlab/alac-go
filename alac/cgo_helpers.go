// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 05 Sep 2017 19:46:15 MSK.
// By https://git.io/c-for-go. DO NOT EDIT.

package alac

/*
#include "alac.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocFileMemory allocates memory for type C.alac_file in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFileMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFileValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfFileValue = unsafe.Sizeof([1]C.alac_file{})

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *File) Ref() *C.alac_file {
	if x == nil {
		return nil
	}
	return x.ref9e6cbf0e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *File) Free() {
	if x != nil && x.allocs9e6cbf0e != nil {
		x.allocs9e6cbf0e.(*cgoAllocMap).Free()
		x.ref9e6cbf0e = nil
	}
}

// NewFileRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewFileRef(ref unsafe.Pointer) *File {
	if ref == nil {
		return nil
	}
	obj := new(File)
	obj.ref9e6cbf0e = (*C.alac_file)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *File) PassRef() (*C.alac_file, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9e6cbf0e != nil {
		return x.ref9e6cbf0e, nil
	}
	mem9e6cbf0e := allocFileMemory(1)
	ref9e6cbf0e := (*C.alac_file)(mem9e6cbf0e)
	allocs9e6cbf0e := new(cgoAllocMap)
	allocs9e6cbf0e.Add(mem9e6cbf0e)

	var cinput_buffer_allocs *cgoAllocMap
	ref9e6cbf0e.input_buffer, cinput_buffer_allocs = (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.InputBuffer)).Data)), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(cinput_buffer_allocs)

	var cinput_buffer_bitaccumulator_allocs *cgoAllocMap
	ref9e6cbf0e.input_buffer_bitaccumulator, cinput_buffer_bitaccumulator_allocs = (C.int)(x.InputBufferBitaccumulator), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(cinput_buffer_bitaccumulator_allocs)

	var csamplesize_allocs *cgoAllocMap
	ref9e6cbf0e.samplesize, csamplesize_allocs = (C.int)(x.Samplesize), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csamplesize_allocs)

	var cnumchannels_allocs *cgoAllocMap
	ref9e6cbf0e.numchannels, cnumchannels_allocs = (C.int)(x.Numchannels), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(cnumchannels_allocs)

	var cbytespersample_allocs *cgoAllocMap
	ref9e6cbf0e.bytespersample, cbytespersample_allocs = (C.int)(x.Bytespersample), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(cbytespersample_allocs)

	var cpredicterror_buffer_a_allocs *cgoAllocMap
	ref9e6cbf0e.predicterror_buffer_a, cpredicterror_buffer_a_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.PredicterrorBufferA)).Data)), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(cpredicterror_buffer_a_allocs)

	var cpredicterror_buffer_b_allocs *cgoAllocMap
	ref9e6cbf0e.predicterror_buffer_b, cpredicterror_buffer_b_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.PredicterrorBufferB)).Data)), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(cpredicterror_buffer_b_allocs)

	var coutputsamples_buffer_a_allocs *cgoAllocMap
	ref9e6cbf0e.outputsamples_buffer_a, coutputsamples_buffer_a_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.OutputsamplesBufferA)).Data)), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(coutputsamples_buffer_a_allocs)

	var coutputsamples_buffer_b_allocs *cgoAllocMap
	ref9e6cbf0e.outputsamples_buffer_b, coutputsamples_buffer_b_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.OutputsamplesBufferB)).Data)), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(coutputsamples_buffer_b_allocs)

	var cuncompressed_bytes_buffer_a_allocs *cgoAllocMap
	ref9e6cbf0e.uncompressed_bytes_buffer_a, cuncompressed_bytes_buffer_a_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.UncompressedBytesBufferA)).Data)), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(cuncompressed_bytes_buffer_a_allocs)

	var cuncompressed_bytes_buffer_b_allocs *cgoAllocMap
	ref9e6cbf0e.uncompressed_bytes_buffer_b, cuncompressed_bytes_buffer_b_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.UncompressedBytesBufferB)).Data)), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(cuncompressed_bytes_buffer_b_allocs)

	var csetinfo_max_samples_per_frame_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_max_samples_per_frame, csetinfo_max_samples_per_frame_allocs = (C.uint32_t)(x.SetinfoMaxSamplesPerFrame), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_max_samples_per_frame_allocs)

	var csetinfo_7a_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_7a, csetinfo_7a_allocs = (C.uint8_t)(x.Setinfo7a), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_7a_allocs)

	var csetinfo_sample_size_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_sample_size, csetinfo_sample_size_allocs = (C.uint8_t)(x.SetinfoSampleSize), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_sample_size_allocs)

	var csetinfo_rice_historymult_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_rice_historymult, csetinfo_rice_historymult_allocs = (C.uint8_t)(x.SetinfoRiceHistorymult), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_rice_historymult_allocs)

	var csetinfo_rice_initialhistory_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_rice_initialhistory, csetinfo_rice_initialhistory_allocs = (C.uint8_t)(x.SetinfoRiceInitialhistory), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_rice_initialhistory_allocs)

	var csetinfo_rice_kmodifier_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_rice_kmodifier, csetinfo_rice_kmodifier_allocs = (C.uint8_t)(x.SetinfoRiceKmodifier), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_rice_kmodifier_allocs)

	var csetinfo_7f_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_7f, csetinfo_7f_allocs = (C.uint8_t)(x.Setinfo7f), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_7f_allocs)

	var csetinfo_80_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_80, csetinfo_80_allocs = (C.uint16_t)(x.Setinfo80), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_80_allocs)

	var csetinfo_82_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_82, csetinfo_82_allocs = (C.uint32_t)(x.Setinfo82), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_82_allocs)

	var csetinfo_86_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_86, csetinfo_86_allocs = (C.uint32_t)(x.Setinfo86), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_86_allocs)

	var csetinfo_8a_rate_allocs *cgoAllocMap
	ref9e6cbf0e.setinfo_8a_rate, csetinfo_8a_rate_allocs = (C.uint32_t)(x.Setinfo8aRate), cgoAllocsUnknown
	allocs9e6cbf0e.Borrow(csetinfo_8a_rate_allocs)

	x.ref9e6cbf0e = ref9e6cbf0e
	x.allocs9e6cbf0e = allocs9e6cbf0e
	return ref9e6cbf0e, allocs9e6cbf0e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x File) PassValue() (C.alac_file, *cgoAllocMap) {
	if x.ref9e6cbf0e != nil {
		return *x.ref9e6cbf0e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *File) Deref() {
	if x.ref9e6cbf0e == nil {
		return
	}
	hxfc4425b := (*sliceHeader)(unsafe.Pointer(&x.InputBuffer))
	hxfc4425b.Data = uintptr(unsafe.Pointer(x.ref9e6cbf0e.input_buffer))
	hxfc4425b.Cap = 0x7fffffff
	// hxfc4425b.Len = ?

	x.InputBufferBitaccumulator = (int32)(x.ref9e6cbf0e.input_buffer_bitaccumulator)
	x.Samplesize = (int32)(x.ref9e6cbf0e.samplesize)
	x.Numchannels = (int32)(x.ref9e6cbf0e.numchannels)
	x.Bytespersample = (int32)(x.ref9e6cbf0e.bytespersample)
	hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&x.PredicterrorBufferA))
	hxf95e7c8.Data = uintptr(unsafe.Pointer(x.ref9e6cbf0e.predicterror_buffer_a))
	hxf95e7c8.Cap = 0x7fffffff
	// hxf95e7c8.Len = ?

	hxff2234b := (*sliceHeader)(unsafe.Pointer(&x.PredicterrorBufferB))
	hxff2234b.Data = uintptr(unsafe.Pointer(x.ref9e6cbf0e.predicterror_buffer_b))
	hxff2234b.Cap = 0x7fffffff
	// hxff2234b.Len = ?

	hxff73280 := (*sliceHeader)(unsafe.Pointer(&x.OutputsamplesBufferA))
	hxff73280.Data = uintptr(unsafe.Pointer(x.ref9e6cbf0e.outputsamples_buffer_a))
	hxff73280.Cap = 0x7fffffff
	// hxff73280.Len = ?

	hxfa9955c := (*sliceHeader)(unsafe.Pointer(&x.OutputsamplesBufferB))
	hxfa9955c.Data = uintptr(unsafe.Pointer(x.ref9e6cbf0e.outputsamples_buffer_b))
	hxfa9955c.Cap = 0x7fffffff
	// hxfa9955c.Len = ?

	hxfa3f05c := (*sliceHeader)(unsafe.Pointer(&x.UncompressedBytesBufferA))
	hxfa3f05c.Data = uintptr(unsafe.Pointer(x.ref9e6cbf0e.uncompressed_bytes_buffer_a))
	hxfa3f05c.Cap = 0x7fffffff
	// hxfa3f05c.Len = ?

	hxf0d18b7 := (*sliceHeader)(unsafe.Pointer(&x.UncompressedBytesBufferB))
	hxf0d18b7.Data = uintptr(unsafe.Pointer(x.ref9e6cbf0e.uncompressed_bytes_buffer_b))
	hxf0d18b7.Cap = 0x7fffffff
	// hxf0d18b7.Len = ?

	x.SetinfoMaxSamplesPerFrame = (uint32)(x.ref9e6cbf0e.setinfo_max_samples_per_frame)
	x.Setinfo7a = (byte)(x.ref9e6cbf0e.setinfo_7a)
	x.SetinfoSampleSize = (byte)(x.ref9e6cbf0e.setinfo_sample_size)
	x.SetinfoRiceHistorymult = (byte)(x.ref9e6cbf0e.setinfo_rice_historymult)
	x.SetinfoRiceInitialhistory = (byte)(x.ref9e6cbf0e.setinfo_rice_initialhistory)
	x.SetinfoRiceKmodifier = (byte)(x.ref9e6cbf0e.setinfo_rice_kmodifier)
	x.Setinfo7f = (byte)(x.ref9e6cbf0e.setinfo_7f)
	x.Setinfo80 = (uint16)(x.ref9e6cbf0e.setinfo_80)
	x.Setinfo82 = (uint32)(x.ref9e6cbf0e.setinfo_82)
	x.Setinfo86 = (uint32)(x.ref9e6cbf0e.setinfo_86)
	x.Setinfo8aRate = (uint32)(x.ref9e6cbf0e.setinfo_8a_rate)
}
