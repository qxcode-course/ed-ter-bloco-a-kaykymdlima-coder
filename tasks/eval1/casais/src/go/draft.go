package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	descasados := make([]int, N)
	var ani int
	contagem := 0
	encontrado := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&ani)
		for k := 0; k < N; k++ {
			if descasados[k] == (-1 * ani) {
				descasados[k] = 0
				contagem++
				encontrado = 1
			}
		}
		if encontrado == 1 {
			encontrado = 0
		} else {
			descasados[i] = ani
		}
	}
	fmt.Printf("%d", contagem)
	fmt.Println()
}
