package main

import "fmt"

func ordem(letras []rune) {
	l := len(letras)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if letras[j] > letras[j+1] {
				letras[j], letras[j+1] = letras[j+1], letras[j]
			}
		}
	}
}
func permutacao(p string) []string {
	letras := []rune(p)
	ordem(letras)
	var resu []string
	lidas := make([]bool, len(letras))
	var atual []rune

	backtrack(letras, lidas, atual, &resu)

	return resu
}

func backtrack(letras []rune, lidas []bool, atual []rune, resu *[]string) {
	if len(atual) == len(letras) {
		*resu = append(*resu, string(atual))
		return
	}

	for i, _ := range letras {
		if lidas[i] {
			continue
		}
		lidas[i] = true
		atual = append(atual, letras[i])
		backtrack(letras, lidas, atual, resu)
		lidas[i] = false
		atual = atual[:len(atual)-1]
	}

}

func main() {
	var entrada string
	fmt.Scanln(&entrada)
	permutacoes := permutacao(entrada)
	for _, saida := range permutacoes {
		fmt.Println(saida)
	}
}
