package main

import "fmt"

func main() {
	var sequenciaori string
	var L int
	fmt.Scan(&sequenciaori, &L)

	sequencia := []byte(sequenciaori)
	solve(sequencia, L, 0)
	fmt.Println(string(sequencia))

}
func solve(sequencia []byte, L, ind int) bool {
	if ind == len(sequencia) {
		return true
	}
	if sequencia[ind] != '.' {
		return solve(sequencia, L, ind+1)
	}
	for i := byte('0'); i <= byte('0'+L); i++ {
		if ehVseq(sequencia, L, ind, i) {
			sequencia[ind] = i
			if solve(sequencia, L, ind+1) {
				return true
			}
		}

	}
	sequencia[ind] = '.'
	return false
}
func ehVseq(sequencia []byte, L, ind int, d byte) bool {
	for i := ind - L; i <= ind+L; i++ {
		if i >= 0 && i < len(sequencia) && i != ind {
			if sequencia[i] == d {
				return false
			}
		}
	}
	return true
}
