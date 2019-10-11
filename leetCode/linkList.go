package leetCode

type Base interface {
	GetName() string
}

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func (l *ListNode) GetNext() *ListNode {
	return l.Next
}

func (l *ListNode) SetNext(node *ListNode) {
	l.Next = node
}
