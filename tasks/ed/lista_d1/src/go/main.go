package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LList struct {
	root *Node
}

func (ll *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	node := ll.root.next
	for node != ll.root {
		sb.WriteString(strconv.Itoa(node.value))
		if node.next != ll.root {
			sb.WriteString(", ")
		}
		node = node.next
	}

	sb.WriteString("]")
	return sb.String()
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{root: root}
}

func (ll *LList) Size() int {
	count := 0
	node := ll.root.next
	for node != ll.root {
		count++
		node = node.next
	}
	return count
}

func (ll *LList) PushFront(value int) {
	newNod := &Node{value: value}
	fi := ll.root.next

	newNod.next = fi
	newNod.prev = ll.root
	ll.root.next = newNod
	fi.prev = newNod
}
func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
}
func (ll *LList) PushBack(value int) {
	newNod := &Node{value: value}
	lt := ll.root.prev
	newNod.prev = lt
	newNod.next = ll.root
	lt.next = newNod
	ll.root.prev = newNod
}
func (ll *LList) PopFront() {
	if ll.root.next == ll.root {
		return
	}
	fi := ll.root.next
	ll.root.next = fi.next
	fi.next.prev = ll.root
}
func (ll *LList) PopBack() {
	if ll.root.prev == ll.root {
		return
	}
	lt := ll.root.prev
	ll.root.prev = lt.prev
	lt.prev.next = ll.root
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
