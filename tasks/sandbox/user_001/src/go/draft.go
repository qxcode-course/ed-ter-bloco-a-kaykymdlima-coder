package main

type Node struct {
	value int
	next  *Node
}

func (l *Lista) __PushBack(node *Node, value int) *Node {

}
func (l *Lista) PushBack(value int) {
	l.head = __PushBack(l.head, value)
}
func main() {

}
