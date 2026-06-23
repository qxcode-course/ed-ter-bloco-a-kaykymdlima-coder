package main

import (
	"fmt"
)

type b struct {
	gaso int
	dist int
}

func main() {
	var n int
	fmt.Scan(&n)
	bs := make([]b, n)
	for i := range bs {
		fmt.Scan(&bs[i].gaso, &bs[i].dist)
	}
	gasoat := 0
	gasoto := 0
	distto := 0

	for i := range bs {
		gasoto += bs[i].gaso
		distto += bs[i].dist
	}
	var st int
	for i := 0; i < n; i++ {
		gasoat += bs[i].gaso - bs[i].dist
		if gasoat < 0 {
			st = i + 1
			gasoat = 0
		}
	}
	fmt.Println(st)
}
