package main

import (
	"bufio"
	"fmt"
	"os"
)

func polonesa(ent string) string {
	var saida []rune
	var pilha []rune
	for _, ch := range ent {
		if ch == ' ' {
			continue
		}
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '^' {
			for len(pilha) > 0 && priori(pilha[len(pilha)-1]) >= priori(ch) {
				top := pilha[len(pilha)-1]
				pilha = pilha[:len(pilha)-1]
				saida = append(saida, top)
				saida = append(saida, ' ')
			}
			pilha = append(pilha, ch)
		} else {
			saida = append(saida, ch)
			saida = append(saida, ' ')
		}
	}
	for len(pilha) > 0 {
		top := pilha[len(pilha)-1]
		pilha = pilha[:len(pilha)-1]
		saida = append(saida, top)
		saida = append(saida, ' ')
	}
	if saida[len(saida)-1] == ' ' {
		saida = saida[:len(saida)-1]
	}
	return string(saida)
}
func priori(operador rune) int {
	if operador == '+' || operador == '-' {
		return 1
	} else if operador == '*' || operador == '/' {
		return 2
	} else if operador == '^' {
		return 3
	}
	return 0
}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	entrada := sc.Text()
	resu := polonesa(entrada)
	fmt.Println(resu)
}
