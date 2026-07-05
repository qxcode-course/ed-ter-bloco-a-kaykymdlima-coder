package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxDetona(bombas [][]int) int {
	nu := len(bombas)
	adjun := make([][]int, nu)
	for i := 0; i < nu; i++ {
		for j := 0; j < nu; j++ {
			if i == j {
				continue
			}
			xi, yi, ri := int64(bombas[i][0]), int64(bombas[i][1]), int64(bombas[i][2])
			xj, yj := int64(bombas[j][0]), int64(bombas[j][1])

			dis := (xi-xj)*(xi-xj) + (yi-yj)*(yi-yj)
			if dis <= ri*ri {
				adjun[i] = append(adjun[i], j)
			}
		}
	}
	maxbomba := 0
	for i := 0; i < nu; i++ {
		visitado := make([]bool, nu)
		cont := 0
		var busca func(node int)
		busca = func(node int) {
			visitado[node] = true
			cont++
			for _, adj := range adjun[node] {
				if !visitado[adj] {
					busca(adj)
				}
			}
		}
		busca(i)
		if cont > maxbomba {
			maxbomba = cont
		}
	}
	return maxbomba
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])

	bombas := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		tok := strings.Fields(scanner.Text())
		row := make([]int, 3)
		for j := 0; j < 3 && j < len(tok); j++ {
			row[j], _ = strconv.Atoi(tok[j])
		}
		bombas[i] = row
	}
	fmt.Println(maxDetona(bombas))
}
