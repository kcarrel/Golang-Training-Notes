func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return prev
}

