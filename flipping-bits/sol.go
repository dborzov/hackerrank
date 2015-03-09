package main

import "fmt"

const MaxInt = ^uint32(0)

func main() {
	var T, M int
	fmt.Scan(&T)
	for i:=0; i<T; i++ {
		fmt.Scan(&M)
		fmt.Println(FlipInt(uint32(M)))
	}
}

func FlipInt(num uint32) uint32{
	return num^MaxInt
}