package main

import "fmt"

func backtrack(lin, n int, col, diap, dias []bool, resu *int) {
	if lin == n {
		*resu++
		return
	}

	for colu := 0; colu < n; colu++ {
		idiap := lin - colu + (n - 1)
		idias := lin + colu
		if col[colu] || diap[idiap] || dias[idias] {
			continue
		}
		col[colu] = true
		diap[idiap] = true
		dias[idias] = true

		backtrack(lin+1, n, col, diap, dias, resu)

		col[colu] = false
		diap[idiap] = false
		dias[idias] = false
	}
}
func main() {
	var n int
	fmt.Scanln(&n)

	col := make([]bool, n)
	diagonalpri := make([]bool, 2*n-1)
	diagonalsec := make([]bool, 2*n-1)
	resu := 0

	backtrack(0, n, col, diagonalpri, diagonalsec, &resu)
	fmt.Println(resu)
}
