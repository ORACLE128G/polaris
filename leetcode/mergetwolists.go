package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

var l1 = ListNode{
	Val:  1,
	Next: nil,
}

var l2 = ListNode{
	Val:  2,
	Next: nil,
}

func main() {

	log.Print(mergeTwoLists(&l1, &l2))
}

// Merge two sorted lists
// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1, l2 *ListNode) (result *ListNode) {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
