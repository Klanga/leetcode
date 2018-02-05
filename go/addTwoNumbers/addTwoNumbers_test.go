package addTwoNumbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}

	if resultListNode := addTwoNumbers(l1, l2); resultListNode.Val != 0 && resultListNode.Next.Val != 0 && resultListNode.Next.Next.Val != 1 {
		t.Errorf("ListNode : '%v' Test Failed!\n", resultListNode)
	} else {
		t.Logf("ListNode : '%v' Test Passed!\n", resultListNode)
	}
}
