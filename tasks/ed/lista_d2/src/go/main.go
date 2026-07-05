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

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root
	return &LList{root: root, size: 0}
}
func (ll *LList) Front() *Node {
	if ll.root.next == ll.root {
		return nil
	}
	return ll.root.next
}
func (ll *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for node := ll.Front(); node != nil; node = node.Next() {
		sb.WriteString(strconv.Itoa(node.Value))
		if node.Next() != nil {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
func (ll *LList) PushFront(value int) {
	newnod := &Node{Value: value, root: ll.root}
	fi := ll.root.next
	newnod.next = fi
	newnod.prev = ll.root

	ll.root.next = newnod
	fi.prev = newnod
	ll.size++
}
func (ll *LList) PushBack(value int) {
	newnod := &Node{Value: value, root: ll.root}
	lt := ll.root.prev
	newnod.prev = lt
	newnod.next = ll.root
	lt.next = newnod
	ll.root.prev = newnod
	ll.size++
}
func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}
func (ll *LList) Back() *Node {
	if ll.size == 0 {
		return nil
	}
	return ll.root.prev
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}
func (ll *LList) Search(value int) *Node {
	for nod := ll.Front(); nod != nil; nod = nod.Next() {
		if nod.Value == value {
			return nod
		}
	}
	return nil
}
func (ll *LList) Insert(node *Node, value int) {
	if node == nil {
		return
	}
	newnod := &Node{Value: value, root: ll.root}
	newnod.prev = node.prev
	newnod.next = node

	node.prev.next = newnod
	node.prev = newnod
	ll.size++
}
func (ll *LList) Remove(node *Node) *Node {
	if node == nil || node == ll.root {
		return nil
	}
	nextnod := node.next
	node.prev.next = node.next
	node.next.prev = node.prev
	ll.size--

	if nextnod == ll.root {
		return nil
	}
	return nextnod
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
			// fmt.Println(ll.Size())
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
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
