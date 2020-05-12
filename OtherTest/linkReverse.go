/*
一个链表，长度为N，从尾部算起，以长度为K分组，分组内倒置。开头第一个分组如果不足K，不需要倒置
*/

package main

import "fmt"

type LinkedNode struct {
	next *LinkedNode
	data int
}

func (myself *LinkedNode) ReverseLinkedList(length int) (nextHead *LinkedNode) {
	var head *LinkedNode = myself
	var p1 *LinkedNode = myself.next
	p2 := head.next
	p3 := head.next
	if p1 == nil {
		return p1
	}
	var count = 1
	for p1.next != nil {
		p2 = head.next
		p3 = p1.next
		head.next = p1.next
		p1.next = p3.next
		p3.next = p2
		count++

		if count == length {
			break
		}
	}
	return p1
}
func (myself *LinkedNode) CreateLinkedList(datas []int) {
	var currentNode = myself
	for _, v := range datas {
		nextNode := new(LinkedNode)
		nextNode.next = currentNode.next
		nextNode.data = v
		currentNode.next = nextNode
	}
}
func (myself *LinkedNode) OutputLinkedList() {
	var c = myself.next
	for c != nil {
		fmt.Printf("%d ", c.data)
		c = c.next
	}
	fmt.Println("")
}
func main() {
	var head = new(LinkedNode)
	head.next = nil
	head.data = 0

	head.CreateLinkedList([]int{2, 3, 4, 5, 6, 7, 8})
	head.OutputLinkedList()

	head.ReverseLinkedList(10)
	head.OutputLinkedList()

	head.ReverseLinkedList(3)
	head.OutputLinkedList()

	currentHead := head
	for currentHead != nil && currentHead.next != nil {
		currentHead = currentHead.ReverseLinkedList(3)
	}
	head.OutputLinkedList()
}
