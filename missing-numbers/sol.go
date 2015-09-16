package main

import (
	"fmt"
	"sort"
)

func main() {
	var Asize int
	fmt.Scanf("%d\n", &Asize)
	A := make([]int, Asize)
	for i := range A {
		fmt.Scanf("%d", &A[i])
	}
	var Bsize int
	fmt.Scanf("%d\n", &Bsize)
	B := make([]int, Bsize)
	for i := range B {
		fmt.Scanf("%d", &B[i])
	}

	solve(A, B)
}

type SortInts []int

func (a SortInts) Len() int           { return len(a) }
func (a SortInts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortInts) Less(i, j int) bool { return a[i] < a[j] }

func solve(A, B []int) {
	mentioned := make(map[int]bool)
	sort.Sort(SortInts(A))
	sort.Sort(SortInts(B))
	// fmt.Printf("Here are the sizes: %v, %v\n", len(A), len(B))

	smaller := A
	bigger := B
	if len(A) >= len(B) {
		smaller = B
		bigger = A
	}

	curSmaller := 0
	for i := range bigger {
		if curSmaller < len(smaller) && bigger[i] == smaller[curSmaller] {
			curSmaller++
		} else {
			if _, ok := mentioned[bigger[i]]; !ok {
				mentioned[bigger[i]] = true
				fmt.Printf("%v ", bigger[i])
			}
		}
	}
}
