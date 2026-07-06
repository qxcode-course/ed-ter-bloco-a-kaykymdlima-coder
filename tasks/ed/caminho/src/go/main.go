package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type NQueue[T any] struct {
	queue *list.List
}

func NNewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		queue: list.New(),
	}
}
func (q *Queue[T]) NEnqueue(value T) {
	q.queue.PushBack(value)
}
func (q *Queue[T]) NDequeue() (T, bool) {
	element := q.queue.Front()
	if element == nil {
		var zero T
		return zero, false
	}
	q.queue.Remove(element)
	value := element.Value.(T)
	return value, true
}

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{p.l - 1, p.c},
		{p.l + 1, p.c},
		{p.l, p.c - 1},
		{p.l, p.c + 1},
	}
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	queue := NNewQueue[Pos]()
	queue.NEnqueue(startPos)
	visited := make(map[Pos]bool)
	par := make(map[Pos]Pos)
	visited[startPos] = true
	for !queue.IsEmpty() {
		atual, _ := queue.NDequeue()
		if atual == endPos {
			break
		}
		for _, n := range atual.getNeig() {
			if match(grid, n, ' ') && !visited[n] {
				visited[n] = true
				par[n] = atual
				queue.Enqueue(n)
			}
		}
	}
	if visited[endPos] {
		atual := endPos
		for atual != startPos {
			grid[atual.l][atual.c] = '.'
			atual = par[atual]
		}
		grid[startPos.l][startPos.c] = '.'
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
