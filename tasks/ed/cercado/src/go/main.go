package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	rs := len(board)
	cs := len(board[0])
	visto := make([][]bool, rs)
	for i := range visto {
		visto[i] = make([]bool, cs)
	}
	var busca func(r, c int)
	busca = func(r, c int) {
		if r < 0 || r >= rs || c < 0 || c >= cs || board[r][c] == 'X' || visto[r][c] {
			return
		}
		visto[r][c] = true
		busca(r+1, c)
		busca(r-1, c)
		busca(r, c+1)
		busca(r, c-1)
	}
	for c := 0; c < cs; c++ {
		busca(0, c)
		busca(rs-1, c)
	}
	for r := 0; r < rs; r++ {
		for c := 0; c < cs; c++ {
			if board[r][c] == 'O' && !visto[r][c] {
				board[r][c] = 'X'
			}
		}
	}

}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
