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

	var a = [2 * N]Pair{}

	for i := 0; i < N; i++ {
		a[2*i] = Pair(Timeses[i][0], true)
		a[2*i+1] = Pair(Timeses[i][1], false)
	}
	sort.Ints(a)

	curMax := 0
	maxFinal := 0

	for i := 0; i < 2*N; i++ {
		if a[i].second {
			curMax++
		} else {
			if curMax > maxFinal {
				maxFinal = curMax
			}
			curMax--
		}
	}

	log.Println(maxFinal)
}
