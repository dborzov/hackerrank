package main

import "fmt"

func main() {
	var N, v int
	fmt.Scanf("%d\n", &N)
	counter := make(map[int]int)
	for i := 0; i < N; i++ {
		fmt.Scanf("%v", &v)
		if _, ok := counter[v]; ok {
			counter[v]++
		} else {
			counter[v] = 1
		}
	}

	for key, val := range counter {
		if val == 1 {
			fmt.Printf("%d\n", key)
		}
	}
}
