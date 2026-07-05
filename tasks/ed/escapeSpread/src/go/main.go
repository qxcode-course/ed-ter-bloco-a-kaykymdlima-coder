package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
const inf = 1_000_000_000_000
const high = 1_000_000_000

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	nl := len(grid)
	nc := len(grid[0])
	count := 0

	var busca func(i, j int)
	busca = func(i, j int) {
		if i < 0 || i >= nl || j < 0 || j >= nc || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		busca(i+1, j)
		busca(i-1, j)
		busca(i, j+1)
		busca(i, j-1)
	}
	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				count++
				busca(i, j)
			}
		}
	}
	return count
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
