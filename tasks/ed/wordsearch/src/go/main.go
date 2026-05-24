package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	linh := len(grid)
	colu := len(grid[0])
	var busca func(int, int, int) bool
	busca = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || j < 0 || i >= linh || j >= colu || grid[i][j] != word[k] {
			return false
		}

		tem := grid[i][j]
		grid[i][j] = '#'

		enco := busca(i+1, j, k+1) || busca(i-1, j, k+1) || busca(i, j-1, k+1) || busca(i, j+1, k+1)
		grid[i][j] = tem
		return enco
	}
	for i := range linh {
		for j := 0; j < colu; j++ {
			if busca(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
