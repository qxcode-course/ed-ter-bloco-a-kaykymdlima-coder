package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nenlaranja(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	type pair struct{ x, y int }
	var queue []pair
	frescas := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, pair{i, j})
			} else if grid[i][j] == 1 {
				frescas++
			}
		}
	}
	if frescas == 0 {
		return 0
	}
	minu := 0
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 && frescas > 0 {
		minu++
		size := len(queue)
		for i := 0; i < size; i++ {
			atual := queue[0]
			queue = queue[1:]
			for _, d := range dirs {
				nr, nc := atual.x+d.x, atual.y+d.y
				if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					frescas--
					queue = append(queue, pair{nr, nc})
				}
			}
		}
	}
	if frescas > 0 {
		return -1
	}
	return minu
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	m, _ := strconv.Atoi(parts[0])
	n, _ := strconv.Atoi(parts[1])
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		tok := strings.Fields(scanner.Text())
		row := make([]int, n)
		for j := 0; j < n; j++ {
			row[j], _ = strconv.Atoi(tok[j])
		}
		grid[i] = row
	}
	fmt.Println(nenlaranja(grid))
}
