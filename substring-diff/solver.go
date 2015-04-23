package main

import (
    "fmt"
    "os"
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
    cur := matcher{}
    suffixer := make(map[string]matcher)


    showMatch := func(m matcher) string{
        return P[:m.Plow] + "["+ P[m.Plow:m.Phigh] + "]"+ P[m.Phigh:] +"-->"+ Q[:m.Qlow] + "[" + Q[m.Qlow:m.Qhigh] + "]" + Q[m.Qhigh:]
    }


    for i:=0; i<= len(P); i++ {
        for j:=0; j<=len(Q); j++ {
            if i==0 || j==0 {
                suffixer[string(i)+","+string(j)] = matcher{
                    Plow:i,
                    Phigh:i,
                    Qlow:j,
                    Qhigh:j,
                }
                continue
            }

            upd := suffixer[string(i-1)+","+string(j-1)]
            if P[i-1]==Q[j-1]{
                upd.Phigh = i
                upd.Qhigh = j
                if cur.Len() < upd.Len() {
                    cur = upd
                }
                suffixer[string(i)+","+string(j)] = upd
                continue
            }


            if upd.s < S {
                upd.s++
                upd.Phigh = i
                upd.Qhigh = j
                if cur.Len() < upd.Len() {
                    cur = upd
                }
                suffixer[string(i)+","+string(j)] = upd
                continue
            }

            if P[upd.Plow] == Q[upd.Qlow] {
                suffixer[string(i)+","+string(j)] = matcher{
                    Plow:i,
                    Phigh:i,
                    Qlow:j,
                    Qhigh:j,
                }
                continue
            }

            upd.Plow++
            upd.Qlow++
            suffixer[string(i)+","+string(j)] = upd


        }
    }
    fmt.Fprintf(os.Stderr, "the optimal is : %v \n", showMatch(cur))
    return cur.Len()
}


type matcher struct{
    Phigh, Plow int
    Qhigh, Qlow int
    s int
}

func (m matcher) Len() int {
    return m.Phigh - m.Plow
}