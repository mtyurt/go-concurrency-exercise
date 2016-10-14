package queue

import (
	"errors"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	length int
}

type IntQueue interface {
	Pop() (int, error)
	Push(int)
	Len() int
}

type NavigatableNode interface {
	Next() *Node
	Prev() *Node
	Value() int
}

func CreateQueue() IntQueue {
	q := &Queue{}

	return q
}

func (q *Queue) Push(value int) {
	if q.Head == nil {
		q.Head = &Node{value: value}
		q.Tail = q.Head
		q.length = 1
	} else {
		newNode := &Node{value: value,
			prev: q.Tail}
		(*q.Tail).next = newNode
		q.Tail = newNode
		q.length = q.length + 1
	}
}

func (q *Queue) Pop() (int, error) {
	if q.Head == nil {
		return 0, errors.New("Queue is empty!")
	}
	value := (*q.Head).value
	q.Head = (*q.Head).next
	if q.Head == nil {
		q.Tail = nil
	}
	q.length = q.length - 1
	return value, nil
}

func (q *Queue) Len() int {
	return q.length
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) GetValue() int {
	return n.value
}
