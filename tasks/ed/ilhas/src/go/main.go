package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	qnt := 0
	for i := range grid {
		for j := range grid {
			if grid[i][j] == '1' {
				qnt++
				explora(grid, i, j)
			}
		}
	}
	return qnt
}
func explora(grid [][]byte, l int, c int) {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) || grid[l][c] == '0' {
		return
	}
	grid[l][c] = '0'
	//esquerda
	explora(grid, l-1, c)
	//direita
	explora(grid, l+1, c)
	//sobe
	explora(grid, l, c+1)
	//descee
	explora(grid, l, c-1)
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
