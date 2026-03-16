package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		n = n * -1
	}
	return n
}
func main() {
	var N int
	fmt.Scan(&N)
	P := make([]int, N)
	P1 := make([]int, N)
	P2 := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&P1[i])
		fmt.Scan(&P2[i])

		if P1[i] < 10 || P2[i] < 10 {
			P[i] = 99
			continue
		}
		P[i] = abs(P1[i] - P2[i])
	}
	ven := 99
	var iven int
	for i := 0; i < N; i++ {
		if i == 0 {
			ven = P[i]
			continue
		}
		if P[i] == ven {
			continue
		} else if P[i] < ven {
			ven = P[i]
			iven = i
		}
	}
	if P[iven] == 99 {
		fmt.Print("sem ganhador\n")
	} else {
		fmt.Printf("%d\n", iven)
	}
}
