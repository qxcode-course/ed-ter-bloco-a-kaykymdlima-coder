package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	rs := len(matrix)
	cs := len(matrix[0])
	me := make([][]int, rs)
	for i := range me {
		me[i] = make([]int, cs)
	}

	pat := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1}}

	var busca func(r, c int) int
	busca = func(r, c int) int {
		if me[r][c] != 0 {
			return me[r][c]
		}
		maxpat := 1
		for _, d := range pat {
			ner, nec := r+d[0], c+d[1]
			if ner >= 0 && ner < rs && nec >= 0 && nec < cs && matrix[ner][nec] > matrix[r][c] {
				path := 1 + busca(ner, nec)
				if path > maxpat {
					maxpat = path
				}
			}
		}
		me[r][c] = maxpat
		return maxpat
	}
	long := 0
	for r := 0; r < rs; r++ {
		for c := 0; c < cs; c++ {
			patl := busca(r, c)
			if patl > long {
				long = patl
			}
		}
	}
	return long
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
