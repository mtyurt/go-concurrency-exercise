package queue

import (
	"testing"
)

func TestPush(t *testing.T) {
	q := Queue{}
	q.Push(5)
	assertNode(t, q.Tail, 5)
	assertNode(t, q.Head, 5)

	q.Push(3)
	assertNode(t, q.Head, 5)
	assertNode(t, q.Tail, 3)

	q.Push(7)
	assertNode(t, q.Head, 5)
	assertNode(t, q.Head.Next(), 3)
	assertNode(t, q.Head.Next().Next(), 7)
	assertNode(t, q.Tail, 7)
}

func TestPop(t *testing.T) {
	q := Queue{}
	q.Push(3)
	q.Push(6)
	q.Push(1)

	assertNode(t, q.Head, 3)
	assertNode(t, q.Tail, 1)

	val, err := q.Pop()
	if err != nil || val != 3 {
		t.Error("Pop failed, expected: ", 3, "actual:", val)
	}
	assertNode(t, q.Head, 6)
	assertNode(t, q.Tail, 1)

	val, err = q.Pop()
	if err != nil || val != 6 {
		t.Error("Pop failed, expected: ", 6, "actual:", val)
	}
	assertNode(t, q.Head, 1)
	assertNode(t, q.Tail, 1)

	val, err = q.Pop()
	if err != nil || val != 1 {
		t.Error("Pop failed, expected: ", 1, "actual:", val)
	}

	if q.Head != nil {
		t.Error("head is not nil")
	}
	if q.Tail != nil {
		t.Error("tail is not nil")
	}
	_, err = q.Pop()
	if err == nil {
		t.Error("error is not returned")
	}
}

func TestLen(t *testing.T) {
	q := Queue{}
	if q.Len() != 0 {
		t.Error("length is wrong")
	}

	q.Push(3)
	if q.Len() != 1 {
		t.Error("length is wrong")
	}
	q.Push(4)
	if q.Len() != 2 {
		t.Error("length is wrong")
	}
	q.Push(5)
	if q.Len() != 3 {
		t.Error("length is wrong")
	}
	q.Pop()
	if q.Len() != 2 {
		t.Error("length is wrong")
	}
	q.Pop()
	if q.Len() != 1 {
		t.Error("length is wrong")
	}
	q.Pop()
	if q.Len() != 0 {
		t.Error("length is wrong")
	}

}

func assertNode(t *testing.T, node *Node, value int) {
	if node == nil {
		t.Error("Node is not set")
	}
	if node.GetValue() != value {
		t.Error("Node Value is wrong, expected:", value, "actual:", node.GetValue())
	}
}
