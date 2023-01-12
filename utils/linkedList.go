package utils

//
// LinkedList stores the Head and Tail of a linked list
//
type LinkedList struct {
	// Head is the beginning node of the list
	Head *Node
	// Tail is the last node of the list
	Tail *Node
}

// NewLinkedList generates a new linked list ob
func NewLinkedList(initNode *Node) *LinkedList {
	return &LinkedList{
		Head: initNode,
		Tail: initNode,
	}
}

// InsertStart adds a new node to the beginning of the list
func (l *LinkedList) InsertStart(node *Node)  {
	node.Next = l.Head
	l.Head = node
}

// InsertEnd adds a new node to the end of the list
func (l *LinkedList) InsertEnd(node *Node) {
	l.Tail.Next = node
	l.Tail = node
}

//
// Node is an element in the doubly linked list with pointers to its neighbours
//
type Node struct {
	// Data is a KV store for node data
	Data map[string]string
	// Next is the pointer to the next node in the list
	Next *Node
}

// NewNode creates a new node object
func NewNode(dat map[string]string) *Node {
	return &Node{Data: dat}
}
