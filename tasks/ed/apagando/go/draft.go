package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N)
	filai := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&filai[i])
	}
	fmt.Scan(&M)
	sairam := make(map[int]bool)
	var ident int
	for i := 0; i < M; i++ {
		fmt.Scan(&ident)
		sairam[ident] = true
	}
	for i := 0; i < N; i++ {
		filaf := filai[i]
		if !(sairam[filaf]) {
			fmt.Printf("%d ", filaf)
		}
	}
	fmt.Println()
}
