package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	x := make([]int, N)
	y := make([]int, N)
	var D string
	fmt.Scan(&D)
	for i := 0; i < N; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	for i := N - 1; i > 0; i-- {
		x[i] = x[i-1]
		y[i] = y[i-1]
	}
	if D == "R" {
		x[0] = x[0] + 1
	} else if D == "L" {
		x[0] = x[0] - 1
	} else if D == "U" {
		y[0] = y[0] - 1
	} else if D == "D" {
		y[0] = y[0] + 1
	}
	for i := 0; i < N; i++ {
		fmt.Printf("%d %d\n", x[i], y[i])
	}
}
