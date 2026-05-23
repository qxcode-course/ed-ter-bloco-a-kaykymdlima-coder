package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

type MultiSet struct {
	data []int
}

func NewMultiSet(valor int) *MultiSet {
	return &MultiSet{data: []int{}}
}
func (k *MultiSet) Busca(valor int) int {
	e := 0
	d := len(k.data)

	for e < d {
		m := (e + d) / 2
		if k.data[m] < valor {
			e = m + 1
		} else {
			d = m
		}
	}
	return e
}
func (k *MultiSet) insert(valor int) {
	in := k.Busca(valor)
	k.data = append(k.data, 0)
	for i := len(k.data) - 1; i > in; i-- {
		k.data[i] = k.data[i-1]
	}
	k.data[in] = valor
}
func (k *MultiSet) show() string {
	return "[" + Join(k.data, ", ") + "]"
}
func (k *MultiSet) erase(valor int) bool {
	in := k.Busca(valor)

	if k.data[in] != valor {
		return false
	}
	for i := in; i < len(k.data)-1; i++ {
		k.data[i] = k.data[i+1]
	}
	k.data = k.data[:len(k.data)-1]
	return true
}
func (k *MultiSet) contains(valor int) bool {
	in := k.Busca(valor)
	return k.data[in] == valor
}
func (k *MultiSet) count(valor int) int {
	qnt := 0
	in := k.Busca(valor)
	for i := in; i < len(k.data); i++ {
		if k.data[i] != valor {
			break
		}
		qnt++
	}
	return qnt
}
func (k *MultiSet) unique() int {
	if len(k.data) == 0 {
		return 0
	}
	qnt := 1
	for i := 1; i < len(k.data); i++ {
		if k.data[i] != k.data[i-1] {
			qnt++
		}
	}
	return qnt
}
func (k *MultiSet) clear() {
	for len(k.data) > 0 {
		k.erase(k.data[0])
	}
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.insert(value)
			}
		case "show":
			fmt.Println(ms.show())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if ms.erase(value) {

			} else {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.count(value))
		case "unique":
			fmt.Println(ms.unique())
		case "clear":
			ms.clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
