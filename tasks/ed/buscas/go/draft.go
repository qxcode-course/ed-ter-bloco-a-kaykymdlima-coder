package main

import "fmt"

func main() {
	var tc, tb int
	fmt.Scan(&tc)
	consulta := make([]string, tc)
	for i := 0; i < tc; i++ {
		fmt.Scan(&consulta[i])
	}
	fmt.Scan(&tb)
	buscas := make([]string, tb)
	for i := 0; i < tb; i++ {
		fmt.Scan(&buscas[i])
	}
	resu := matchingStrings(buscas, consulta)
	for i, qnt := range resu {
		if i == len(resu)-1 {
			fmt.Printf("%d", qnt)
			break
		}
		fmt.Printf("%d ", qnt)
	}
	fmt.Println()
}
func matchingStrings(buscas []string, consulta []string) []int {
	matchingqnt := make(map[string]int)
	for _, stri := range consulta {
		matchingqnt[stri]++
	}
	resu := make([]int, len(buscas))
	for i, stri := range buscas {
		resu[i] = matchingqnt[stri]
	}
	return resu
}
