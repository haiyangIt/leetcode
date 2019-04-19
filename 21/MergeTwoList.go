/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddToNewList(prev, p *ListNode) (*ListNode, *ListNode) {
	prev.Next = p
	p = p.Next
	prev = p
	return prev, p
}

func MergeTwoSortList(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1 == nil && l2 != nil {
		return l2
	}

	head := new(ListNode)
	prev := head
	p1 := l1
	p2 := l2

	for p1 != nil && p2 != nil {
		if p1.Val >= p2.Val {
			prev, p2 = AddToNewList(prev, p2)
		} else {
			prev, p1 = AddToNewList(prev, p1)
		}
	}

	for p1 != nil {
		prev, p1 = AddToNewList(prev, p1)
	}
	for p2 != nil {
		prev, p2 = AddToNewList(prev, p2)
	}

	return head.Next
}

func main() {

}
