package main

import (
	"log"
)

type Times struct {
	Start, End int
}

func main() {
	var Timeses = []Times{{1200, 1245}, {1200, 1400}, {1730, 1800}}
	var maxauto = 0
	for m, _ := range Timeses {

		for j := Timeses[m].Start; j < Timeses[m].End; j++ {

			for i := Timeses[m+1].Start; i < Timeses[m+1].End; i++ {

				if j == i {

					break
				} else {
					j = j + 1
					i = i + 1
				}
			}
		}
		maxauto = maxauto + 1

		log.Println(maxauto)
	}

}
