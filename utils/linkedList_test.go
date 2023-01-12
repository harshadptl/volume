package utils

import "testing"

func TestNewLinkedList(t *testing.T) {
	n := Node{
		Data: map[string]string{"key": "value"},
		Next: nil,
	}

	l := NewLinkedList(&n)
	if l.Head != &n {
		t.Log("incorrect Head pointer")
		t.Fail()
	}
	if l.Tail != &n {
		t.Log("incorrect Tail pointer")
		t.Fail()
	}
}

func TestLinkedList_InsertStart(t *testing.T) {
	n := Node{
		Data: map[string]string{"key": "value"},
		Next: nil,
	}
	l := NewLinkedList(&n)

	n1 := Node{
		Data: nil,
		Next: nil,
	}
	l.InsertStart(&n1)
	if l.Head != &n1 {
		t.Log("incorrect Head pointer")
		t.Fail()
	}
	if l.Tail != &n {
		t.Log("incorrect Tail pointer")
		t.Fail()
	}
	if l.Head.Next != &n {
		t.Log("incorrect Head.Next pointer")
		t.Fail()
	}
}

func TestLinkedList_InsertEnd(t *testing.T) {
	n := Node{
		Data: map[string]string{"key": "value"},
		Next: nil,
	}
	l := NewLinkedList(&n)

	n1 := Node{
		Data: nil,
		Next: nil,
	}
	l.InsertEnd(&n1)
	if l.Head != &n {
		t.Log("incorrect Head pointer")
		t.Fail()
	}
	if l.Tail != &n1 {
		t.Log("incorrect Tail pointer")
		t.Fail()
	}
	if l.Head.Next != &n1 {
		t.Log("incorrect Head.Next pointer")
		t.Fail()
	}
}

func TestNewNode(t *testing.T) {
	n := NewNode(
		map[string]string{
			"a": "1",
			"b": "2",
		})
	if n.Data["a"] != "1" {
		t.Log("incorrect Data")
		t.Fail()
	}
	if n.Data["b"] != "2" {
		t.Log("incorrect Data")
		t.Fail()
	}
	if n.Next != nil {
		t.Log("incorrect Next pointer")
		t.Fail()
	}
}