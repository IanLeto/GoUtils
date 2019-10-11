package leetCode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var countM = 0
	prev := &ListNode{}
loop:
	for head.Next != nil {
		head = head.Next
		countM += 1
		if countM == m {
			//temp := head.Next
			head.Next = prev
			prev = head
			//head = temp
		}
		if countM == n {
			break loop
		}
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	if head == nil || head.Next == nil {
		return head
	}
	for head.Next != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	head.Next = prev
	return head
}
