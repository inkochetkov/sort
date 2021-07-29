package main

import (
	"log"
	"sort"
)

type Pair struct {
	first  int
	second bool
}

func main() {
	var Timeses = [][2]int{{1200, 1245}, {1200, 1400}, {1730, 1800}}
	N := len(Timeses)
	answer := maxauto(Timeses, N)
	log.Println(answer)
}

func maxauto(Timeses, N int) (answer int) {

	a := Pair[2*N]

	for i := 0; i < N; i++ {
		a[2*i] = Pair(Timeses[i][0], true)
		a[2*i+1] = Pair(Timeses[i][1], false)
	}
	sort.Ints(a)

	curMax := 0
	maxFinal := 0

	for i := 0; i < 2*N; i++ {
		if a[i][1] {
			curMax = curMax + 1
		} else {
			if curMax > maxFinal {
				maxFinal = curMax
			}
			curMax = curMax - 1
		}
	}
	answer = maxFinal
	return answer
}
