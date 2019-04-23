/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddToNewList(prev, p *ListNode) (*ListNode, *ListNode) {
	prev.Next = p
	prev = p
	p = p.Next
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
			prevN, p2N := AddToNewList(prev, p2)
			prev = prevN
			p2 = p2N
		} else {
			prevN, p1N := AddToNewList(prev, p1)
			prev = prevN
			p1 = p1N
		}
	}

	for p1 != nil {
		prevN, p1N := AddToNewList(prev, p1)
		prev = prevN
		p1 = p1N
	}
	for p2 != nil {
		prevN, p2N := AddToNewList(prev, p2)
		prev = prevN
		p2 = p2N
	}

	return head.Next
}

func CreateListNodes(datas []int) *ListNode {
	var head *ListNode = new(ListNode)
	prev := head
	for _, d := range datas {
		newNode := CreateListNode(d)
		prev.Next = newNode
		prev = newNode
	}
	return head.Next
}

func CreateListNode(data int) *ListNode {
	l := new(ListNode)
	l.Val = data
	l.Next = nil
	return l
}

func PrintLinkList(l1 *ListNode) {
	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
}

func main() {
	l1 := []int{1, 2, 3, 4, 5, 8, 9}
	l2 := []int{1, 3, 5, 7, 9}

	l1L := CreateListNodes(l1)
	PrintLinkList(l1L)
	l2L := CreateListNodes(l2)
	PrintLinkList(l2L)

	r := MergeTwoSortList(l1L, l2L)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
