package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (q *Queue[T]) Size() int {
	return q.items.Len()
}

func main() {
	//instancia
	sc := bufio.NewScanner(os.Stdin)
	q := NewQueue[string]()
	//pega os times e coloca na fila
	for time := 'A'; time <= 'P'; time++ {
		q.Enqueue(string(time))
	}
	//para cada elemento da fila
	for q.Size() > 1 {
		//para
		if !sc.Scan() {
			break
		}
		//
		line := sc.Text()
		//separa as entrada
		args := strings.Fields(line)
		//pega a entrada a  esqueda e direita
		golx, _ := strconv.Atoi(args[0])
		goly, _ := strconv.Atoi(args[1])
		//remove os times a serem usados
		timex := q.Dequeue()
		timey := q.Dequeue()
		//retorna pra fila o time vencedor
		if golx > goly {
			q.Enqueue(timex)
		} else {
			q.Enqueue(timey)
		}
	}
	fmt.Println(q.Dequeue())
}
