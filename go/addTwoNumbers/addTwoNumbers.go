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
	sum := 0

	for p != nil || q != nil {
		sum /= 10

		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}

		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
	}

	if sum/10 != 0 {
		curr.Next = &ListNode{
			sum / 10,
			nil,
		}
	}
	return node.Next
}
