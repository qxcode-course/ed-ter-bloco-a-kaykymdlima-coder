package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}
func equals(lla, llb *LList) bool {
	nua := lla.root.next
	nub := llb.root.next

	for nua != lla.root && nub != llb.root {
		if nua.Value != nub.Value {
			return false
		}
		nua = nua.next
		nub = nub.next
	}
	return nua == lla.root && nub == llb.root
}

func addsorted(l *LList, value int) {
	n := l.root.next
	for n != l.root && n.Value < value {
		n = n.next
	}
	l.insertBefore(n, value)
}
func reverse(l *LList) {
	atual := l.root
	for {
		atual.next, atual.prev = atual.prev, atual.next
		atual = atual.prev
		if atual == l.root {
			break
		}
	}
}
func merge(lla, llb *LList) *LList {
	nua := lla.root.next
	nub := llb.root.next

	for nub != llb.root {
		if nua != lla.root && nua.Value <= nub.Value {
			nua = nua.next
		} else {
			llb.insertBefore(nua, nub.Value)
			nub = nub.next
		}
	}
	return lla
}
func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func (l *LList) String() string {
	valuel := []string{}

	for n := l.root.next; n != l.root; n = n.next {
		valuel = append(valuel, strconv.Itoa(n.Value))
	}
	return "[" + strings.Join(valuel, ", ") + "]"
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
