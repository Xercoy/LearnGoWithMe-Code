package main

import (
	"fmt"
	"unsafe"
)

func main() {
	intArray := [...]int{1, 2}

	fmt.Printf("\nintArray: %v\n", intArray)

	intPtr := &intArray[0]
	fmt.Printf("\nintPtr=%p, *intPtr=%d.\n", intPtr, *intPtr)

	addressHolder := uintptr(unsafe.Pointer(intPtr)) + unsafe.Sizeof(intArray[0])

	intPtr = (*int)(unsafe.Pointer(addressHolder))

	fmt.Printf("\nintPtr=%p, *intPtr=%d.\n\n", intPtr, *intPtr)
}
