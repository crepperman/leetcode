func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := newNode(0)
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		cur.Next = newNode(sum)
		cur = cur.Next
	}

	return dummy.Next
}

var nodePool = sync.Pool{
	New: func() any {
		return &ListNode{}
	},
}

func newNode(val int) *ListNode {
	n := nodePool.Get().(*ListNode)
	n.Val = val
	n.Next = nil
	return n
}
