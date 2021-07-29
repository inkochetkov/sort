package main

import (
	"log"
	"sort"
)

type Pair struct {
	First  int
	Second bool
}

func main() {
	var Timeses = [][2]int{{1200, 1300}, {1230, 1500}, {1600, 1700}}
	N := len(Timeses)

	var A = make([]Pair, 6)

	for i := 0; i < N; i++ {
		A[2*i] = Pair{Timeses[i][0], true}
		A[2*i+1] = Pair{Timeses[i][1], false}
		//	log.Println(Timeses[i][0], Timeses[i][1])

	}
	//log.Println(A)
	sort.Slice(A, func(i, j int) bool { return A[i].First < A[j].First })

	curMax := 0
	maxFinal := 0

	for i := 0; i < 2*N; i++ {
		if A[i].Second {
			curMax = curMax + 1
		} else {
			if curMax > maxFinal {
				maxFinal = curMax
			}
			curMax = curMax - 1
		}
	}

	log.Println(maxFinal)
}
