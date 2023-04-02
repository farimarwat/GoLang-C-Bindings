package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lexample
#include "example.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	num1 := 1
	num2 := 2

	cnum1 := C.int(num1)
	cnum2 := C.int(num2)

	defer C.free(unsafe.Pointer(&cnum1))
	defer C.free(unsafe.Pointer(&cnum2))
	result := C.sum(cnum1, cnum2)

	fmt.Printf("Result is: %v\n", int(result))
}
