package main

import (
	"fmt"
	"testing"
)

func TestXorIntsAtPos(t *testing.T) {
	fmt.Printf("gets us %v\n", findMaxXor(10, 15))
	fmt.Printf("ein %v\n", xorIntsAtPos(10, 15, 2))
	fmt.Printf("zwei %v\n", xorIntsAtPos(10, 15, 1))
	fmt.Printf("drei %v\n", xorIntsAtPos(10, 15, 0))
	fmt.Printf("xor %v %v at pos %v is %v\n", 8, 4, 2, xorIntsAtPos(8, 4, 2))
}
