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
	for i := 1; i <= N; i++ {
		fmt.Scan(&P1[i])
		fmt.Scan(&P2[i])
	}
	for j := 1; j <= N; j++ {
		if P1[j] < 10 || P2[j] < 10 {
			P[j] = 99
			continue
		}
		P[j] = abs(P1[j] - P2[j])
	}
	var ven int
	var iven int
	for i := 1; i < N; i++ {
		if P[i] == 1 {
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
		fmt.Print("sem ganhador")
	} else {
		fmt.Printf("%d", iven)
	}
}
