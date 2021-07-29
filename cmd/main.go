package main

import (
	"log"
)

type Pair struct {
	first  int
	second bool
}

func main() {
	var Timeses = [][2]int{{1200, 1300}, {1230, 1500}, {1245, 1700}}
	N := len(Timeses)

	var A = [6]Pair{}

	for i := 0; i < N; i++ {
		A[2*i] = Pair{Timeses[i][0], true}
		A[2*i+1] = Pair{Timeses[i][1], false}
		//	log.Println(Timeses[i][0], Timeses[i][1])

	}
	//log.Println(A)

	curMax := 0
	maxFinal := 0

	for i := 0; i < 2*N; i++ {
		if A[i].second {
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
