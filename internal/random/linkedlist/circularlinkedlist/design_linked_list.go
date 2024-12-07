package main

// 707. 设计链表
// https://leetcode.cn/problems/design-linked-list

type MyListNode struct {
	Val  int
	Prev *MyListNode
	Next *MyListNode
}

type MyLinkedList struct {
	dummyHead *MyListNode
	size      int
}

func Constructor() MyLinkedList {
	head := &MyListNode{}
	head.Prev = head
	head.Next = head

	return MyLinkedList{
		dummyHead: head,
		size:      0,
	}
}

func (list *MyLinkedList) Get(index int) int {
	if index >= list.size {
		return -1
	}

	curNode := list.dummyHead.Next
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}

	return curNode.Val
}

func (list *MyLinkedList) AddAtHead(val int) {
	dummyHead := list.dummyHead
	originalHead := dummyHead.Next

	newNode := &MyListNode{
		Val:  val,
		Prev: dummyHead,
		Next: originalHead,
	}

	originalHead.Prev = newNode
	dummyHead.Next = newNode
	list.size++
}

func (list *MyLinkedList) AddAtTail(val int) {
	dummyHead := list.dummyHead
	tailNode := dummyHead.Prev

	newNode := &MyListNode{
		Val:  val,
		Prev: tailNode,
		Next: dummyHead,
	}

	dummyHead.Prev = newNode
	tailNode.Next = newNode
	list.size++

	// 等价于 AddAtIndex(list.size)
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index > list.size {
		return
	}

	curNode := list.dummyHead
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}

	nextNode := curNode.Next

	newNode := &MyListNode{
		Val:  val,
		Prev: curNode,
		Next: nextNode,
	}
	curNode.Next = newNode
	nextNode.Prev = newNode
	list.size++
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index >= list.size {
		return
	}

	curNode := list.dummyHead
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}

	needDelNode := curNode.Next

	nextNode := needDelNode.Next
	curNode.Next = nextNode
	nextNode.Prev = curNode
	list.size--
}
