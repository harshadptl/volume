package utils

// DoublyLinkedList stores the Head and Tail of a doubly linked list
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// NewLinkedList generates a new linked list ob
func NewLinkedList(initNode *Node) *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: initNode,
		Tail: initNode,
	}
}

// InsertStart adds a new node to the beginning of the list
func (l *DoublyLinkedList) InsertStart(node *Node)  {
	node.Next = l.Head
	l.Head.Prev = node
	l.Head = node
}

// InsertEnd adds a new node to the end of the list
func (l *DoublyLinkedList) InsertEnd(node *Node) {
	node.Prev = l.Tail
	l.Tail.Next = node
	l.Tail = node
}

// Node is an element in the doubly linked list with pointers to its neighbours
type Node struct {
	Data map[string]string

	Next *Node
	Prev *Node
}

// NewNode creates a new node object
func NewNode(dat map[string]string) *Node {
	return &Node{Data: dat}
}
