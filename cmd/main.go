package main

import (
	"log"
)

type Times struct {
	Start, End int
}

func main() {
	var calls = []Times{{1200, 1245}, {1200, 1400}, {1230, 1400}}
	for _, t := range calls {
		log.Println(t.Start, t.End)
	}

}
