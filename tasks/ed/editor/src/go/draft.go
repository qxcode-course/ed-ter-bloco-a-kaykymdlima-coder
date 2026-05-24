package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	text := scan.Text()
	e := []rune{}
	d := []rune{}

	for _, cmd := range text {
		switch cmd {
		case 'R':
			e = append(e, '\n')
		case 'B':
			if len(e) > 0 {
				e = e[:len(e)-1]
			}
		case 'D':
			if len(d) > 0 {
				d = d[:len(d)-1]
			}
		case '<':
			if len(e) > 0 {
				ultimo := e[len(e)-1]
				e = e[:len(e)-1]
				d = append(d, ultimo)
			}
		case '>':
			if len(d) > 0 {
				ultimo := d[len(d)-1]
				d = d[:len(d)-1]
				e = append(e, ultimo)
			}
		default:
			e = append(e, cmd)
		}

	}
	for _, cmd := range e {
		fmt.Print(string(cmd))
	}
	fmt.Print("|")
	for i := len(d) - 1; i >= 0; i-- {
		fmt.Print(string(d[i]))
	}
	fmt.Println()
}
