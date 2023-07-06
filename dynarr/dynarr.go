package dynarr

import (
	"fmt"
	"unsafe"
)

type Int int32

type DynArray struct {
	ptr unsafe.Pointer
	cap int
	len int
}

func NewDynArray() *DynArray {
	return &DynArray{
		ptr: nil,
		cap: 0,
		len: 0,
	}
}

func (a *DynArray) Append(element Int) {
	if a.len >= a.cap {
		newCap := (a.cap + 1) * 2
		newPtr := malloc(newCap * sizeof(Int(0)))
		memcopy(a.ptr, newPtr, a.cap*sizeof(Int(0)))
		a.ptr = newPtr
		a.cap = newCap
	}
	writer := (*Int)(unsafe.Add(a.ptr, a.len*sizeof(Int(0))))
	*writer = element
	a.len = a.len + 1
}

func (a *DynArray) Get(index int) (Int, error) {
	if index > a.len {
		return 0, fmt.Errorf("index out of range")
	}
	reader := (*Int)(unsafe.Add(a.ptr, index*sizeof(Int(0))))
	return *reader, nil
}

func (a *DynArray) Size() int {
	return a.len
}

func sizeof(x interface{}) int {
	return int(unsafe.Sizeof(x))
}

func malloc(size int) unsafe.Pointer {
	addr := &make([]byte, size)[0]
	return unsafe.Pointer(addr)
}

func memcopy(src unsafe.Pointer, dest unsafe.Pointer, numOfBytes int) {
	for i := 0; i < numOfBytes; i++ {
		s := (*byte)(src)
		d := (*byte)(dest)
		*d = *s
		src = unsafe.Add(src, 1)
		dest = unsafe.Add(dest, 1)
	}
}
