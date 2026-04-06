package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	return "[" + tostrrecu(vet) + "]"
}
func tostrrecu(vet []int) string {
	//ultimo caso
	if len(vet) == 1 {
		return fmt.Sprintf("%d", vet[0])
	}
	//outros casos
	return fmt.Sprintf("%d, ", vet[0]) + tostrrecu(vet[1:])
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	return "[" + tostrrevrecu(vet) + "]"
}
func tostrrevrecu(vet []int) string {
	if len(vet) == 1 {
		return fmt.Sprintf("%d", vet[0])
	}
	//msm coisa, soq pega o id do ultimo caso
	ultid := len(vet) - 1
	return fmt.Sprintf("%d, ", vet[ultid]) + tostrrevrecu(vet[:ultid])
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	//quando o espaçamento acaba, o vetor tbm
	if len(vet) <= 1 {
		return
	}
	//troca as pontas
	vetori := vet[0]
	vet[0] = vet[len(vet)-1]
	vet[len(vet)-1] = vetori
	///reduz o espacamento do vetor a cada giro
	reverse(vet[1 : len(vet)-1])
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	//se 0, 0
	if len(vet) == 0 {
		return 0
	}
	//faz a soma recu
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	//se 0, 0
	if len(vet) == 0 {
		return 1
	}
	//faz a msm coisa da soma recu soq *
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	// if 0 -1
	if len(vet) == 0 {
		return -1
	}
	var rec func(v []int) (int, int)
	//recursiva interna
	rec = func(v []int) (int, int) {
		// se so 1
		if len(v) == 1 {
			return 0, v[0]
		}
		//recursiva
		idres, valres := rec(v[1:])
		idslide := idres + 1

		///compara
		if v[0] <= valres {
			return 0, v[0]
		} else {
			return idslide, valres
		}
	}
	indiceFinal, _ := rec(vet)
	return indiceFinal
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
