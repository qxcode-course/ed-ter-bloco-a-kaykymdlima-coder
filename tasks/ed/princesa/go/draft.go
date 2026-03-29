package main

import "fmt"

func main() {
	var N, E int
	fmt.Scan(&N, &E)

	vivos := make([]int, N)
	espada := E - 1
	tvivos := N
	for i := 0; i < N; i++ {
		vivos[i] = i + 1
	}
	for tvivos > 0 {
		fmt.Print("[ ")
		for i := 0; i < N; i++ {
			if vivos[i] != 0 {
				if i == espada {
					fmt.Printf("%d> ", i)
				} else {
					fmt.Printf("%d ", i)
				}
			}
		}
		fmt.Println("]")
		if tvivos == 1 {
			break
		}
		for i := 0; i < N; i++ {
			if vivos[i] == 0 {
				continue
			}
			if i == espada {
				if i == N-1 {

				}
				vivos[i+1] = 0
				espada
			}
		}

	}

}
