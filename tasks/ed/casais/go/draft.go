package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	descasados := make([]int, N)
	var animal int
	var cont = 0
	var v int
	for i := 0; i < N; i++ {
		fmt.Scan(&animal)
		for j := 0; j < N; j++ {
			if descasados[j] == (animal * -1) {
				descasados[j] = 0
				cont = cont + 1
				v = 1
				break
			}
		}
		if v == 1 {
			v = 0
			continue
		} else {
			descasados[i] = animal
			v = 0
		}
	}
	fmt.Printf("%d\n", cont)
}
