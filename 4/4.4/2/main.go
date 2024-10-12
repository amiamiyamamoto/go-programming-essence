package main

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	ptr := C.malloc(C.size_t(100)) // Cのmallocを使ってメモリを確保
	C.free(unsafe.Pointer(ptr))    // Cのfreeを呼び出してメモリを解放
}
