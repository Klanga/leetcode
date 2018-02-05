package addTwoNumbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{
		Val: 0,
	}

	curr := node

	p := l1
	q := l2
	carry := 0

	for p != nil || q != nil {
		x := 0
		y := 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}

		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return node.Next
}
