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
	if l1 == nil && l2 == nil {
		return nil
	}

	node := &ListNode{}
	if l1 != nil && l2 == nil {
		node.Val = l1.Val
		if node.Val >= 10 {
			node.Val %= 10
			if l1.Next != nil {
				l1.Next.Val++
			} else {
				node.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
				return node
			}
		}
		node.Next = addTwoNumbers(l1.Next, nil)
	} else if l2 != nil && l1 == nil {
		node.Val = l2.Val
		if node.Val >= 10 {
			node.Val %= 10
			if l2.Next != nil {
				l2.Next.Val++
			} else {
				node.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
				return node
			}
		}
		node.Next = addTwoNumbers(nil, l2.Next)
	} else {
		node.Val = l1.Val + l2.Val
		if node.Val >= 10 {
			node.Val %= 10
			if l1.Next != nil {
				l1.Next.Val++
			} else if l2.Next != nil {
				l2.Next.Val++
			} else {
				node.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
				return node
			}
		}
		node.Next = addTwoNumbers(l1.Next, l2.Next)
	}

	return node
}
