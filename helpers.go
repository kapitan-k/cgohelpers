package cgohelpers

// #include <stdlib.h>
// #include "helpers.h"
import "C"

import (
	"reflect"
	"unsafe"
)

// if c == 0 returns false
// for all other values true
func CharToBool(c C.char) bool {
	return c != 0
}

func BoolToChar(v bool) C.char {
	if v {
		return 1
	}

	return 0
}

// if c == 0 returns false
// for all other values true
func IntToBool(c C.int) bool {
	return c != 0
}

func BoolToInt(v bool) C.int {
	if v {
		return 1
	}

	return 0
}

// it is used to load data with a vector which is represented by ptrs
type MultiData struct {
	// How many ptrs are set
	Cnt   uint64
	Ptrs  []unsafe.Pointer
	Sizes []uint64
}

func (self *MultiData) Free() {
	ptrs := self.Ptrs[:self.Cnt]
	CFreeUnsafeMulti(ptrs)
	CFreeUnsafe(unsafe.Pointer(&self.Ptrs[0]))
}

// it is used to load data in a predefined (in go allocated) ptrs slice
type ExternalMultiData struct {
	MultiData
	// If more are expected to come
	// -> if the corresponding C function is to be invoked again
	HasMore bool
}

func ExternalMultiDataCreate(cap int) (self ExternalMultiData) {
	self.Ptrs = make([]unsafe.Pointer, cap)
	self.Sizes = make([]uint64, cap)
	return
}

func (self *ExternalMultiData) Free() {
	ptrs := self.Ptrs[:self.Cnt]
	CFreeUnsafeMulti(ptrs)
}

func CFreeUnsafe(ptr unsafe.Pointer) {
	C.free(ptr)
}

func CFreeUnsafeMulti(ptrs []unsafe.Pointer) {
	C.free_multi((*unsafe.Pointer)(unsafe.Pointer(&ptrs[0])), C.size_t(len(ptrs)))
}

func CFreeBuf(data []byte) {
	C.free(unsafe.Pointer(&data[0]))
}

func CFreeBufMulti(datas [][]byte) {
	l := len(datas)
	ptrs := make([]unsafe.Pointer, l)
	for i, data := range datas {
		ptrs[i] = unsafe.Pointer(&data[0])
	}
	CFreeUnsafeMulti(ptrs)
}

func SetBytesSliceHeader(pData *[]byte, ptr uintptr, len, cap int) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(pData))
	sh.Data = ptr
	sh.Len = len
	sh.Cap = cap
}
