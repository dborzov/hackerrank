package main
import "fmt"

func main() {
    var N int
    var counter = make(map[int]int)
    fmt.Scanf("%d\n",&N)
    A:= make([]int, N)
    for i:=0; i<N; i++ {
        fmt.Scanf("%v",&A[i])
        if _, ok := counter[A[i]]; ok {
        	counter[A[i]]++
        } else {
        	counter[A[i]]= 1
        } 
    }
    // fmt.Printf("A %#v \n", A) 
    // fmt.Printf("stats %#v \n", counter)
    for k, v:= range counter {
    	if v== 1 {
    		fmt.Println(uint(k))
    		break
    	}
    }
}