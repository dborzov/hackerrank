package main

import "fmt"

func main() {
	var testNum int
	fmt.Scanf("%d\n", &testNum)
	for i := 0; i < testNum; i++ {
		runTest()
	}
}

func runTest() {
	var arLen int
	fmt.Scanf("%d\n", &arLen)
	arr := make([]int, arLen)
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}
	if out := solver(arr); out {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func solver(array []int) bool {
	var sum int
	for _, el := range array {
		sum += el
	}
	var subsum int
	for _, el := range array {
		if 2*subsum+el == sum {
			return true
		}
		subsum += el
	}
	return false
}
