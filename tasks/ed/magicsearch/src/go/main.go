package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	ini := 0
	fim := len(slice)

	for ini < fim {
		meio := (ini + fim) / 2
		pos := meio
		if slice[meio] == value {
			for i := meio + 1; i < fim; i++ {
				if slice[i] == value {
					pos = i
				} else {
					break
				}
			}
			return pos
		} else if slice[meio] < value {
			ini = meio + 1
		} else if slice[meio] > value {
			fim = meio
		}
	}
	return ini
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
