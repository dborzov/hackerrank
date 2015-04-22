package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	var input1, input2 string
	var S int
	for i := 0; i < num; i++ {
		fmt.Scanf("%d%s%s\n", &S, &input1, &input2)
		L := FindMaxL(S, input1, input2)
		fmt.Printf("%v\n", L)
	}

}

func FindMaxL(S int, P, Q string) int {
	fmt.Fprintf(os.Stderr, strings.Repeat("~", 10)+"\n")
	defer fmt.Fprintf(os.Stderr, strings.Repeat("-", 10)+"\n")
    fmt.Fprintf(os.Stderr, "   Executed with S:%v,\n     P: \"%v\"\n     Q: \"%v\" \n", S, P, Q)

	A := make([][]match, len(P))
	for i := 0; i < len(P); i++ {
		A[i] = make([]match, len(Q))
	}

	for i := 0; i < len(P); i++ {
		for j := 0; j < len(Q); j++ {
			if Get(A,i-1, j-1).Subsequents(i, j) && P[i] == Q[j] {
				A[i][j] = Get(A,i-1, j-1)
				A[i][j].indices = append(A[i][j].indices, i, j)
				continue
			}

			A[i][j] = Get(A,i-1, j)
			if len(Get(A,i, j-1).indices) > len(Get(A,i-1, j).indices) {
				A[i][j] = Get(A,i, j-1)
			}
		}
	}

    fmt.Fprintf(os.Stderr, "    [%v]\n    ---> %v <---\n", A[len(P)-1][len(Q)-1].indices,A[len(P)-1][len(Q)-1].String(P))

	return 42
}

type match struct {
	indices []int
	s       int
}

func (m match) Subsequents(i, j int) bool {
	if len(m.indices) < 2 {
		return true
	}
	return m.indices[len(m.indices)-2] == i-1 && m.indices[len(m.indices)-1] == j-1
}

func (m match) String(P string) string{
    var out []byte
    for i:=0; i< len(m.indices); i +=2 {
        out = append(out, P[m.indices[i]])
    }
    return string(out)
}


func Get(A [][]match, i, j int) match {
	if i < 0 || j < 0 {
		return match{}
	}
	return A[i][j]
}
