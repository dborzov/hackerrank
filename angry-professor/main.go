package main

import "fmt"

func main() {
	var testsNum int
	fmt.Scanf("%d\n", &testsNum)
	for t := 0; t < testsNum; t++ {
		runTest()
	}
}

func runTest() {
	var N, K, v int
	count := 0
	fmt.Scanf("%d%d\n", &N, &K)
	for n := 0; n < N; n++ {
		fmt.Scanf("%d", &v)
		if v <= 0 {
			count++
		}
	}
	if count >= K {
		fmt.Printf("NO\n")
	} else {
		fmt.Printf("YES\n")
	}

}
