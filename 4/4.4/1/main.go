package main

/*
#include <stdio.h>
#include <mylib.h>

// #cgo LDFLAGS: -L. -lmylib
// #cgo CFLAGS: -I.
*/
import "C"
import "unsafe"

func main() {
	s := "index.dat"

	ptr := C.CString(s)
	defer C.free(unsafe.Pointer(ptr))
	C.mylib_update_data(ptr)
}
