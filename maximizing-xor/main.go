package main

import "fmt"

func main() {
	var L, R int
	fmt.Scanf("%v\n%v\n", &L, &R)
	fmt.Printf("%v\n", findMaxXor(L, R))
}

func findMaxXor(L, R int) int {
	var h uint
	for ord := uint(10); ord > 0; ord-- {
		if xorIntsAtPos(L, R, ord) {
			h = ord + 1
			break
		}
	}
	return 1<<h - 1

}

func xorIntsAtPos(L, R int, ord uint) bool {
	l := L & (1 << uint(ord))
	r := R & (1 << uint(ord))
	return l^r != 0
}
