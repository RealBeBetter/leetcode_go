package main

// 707. 设计链表
// https://leetcode.cn/problems/design-linked-list

type MyListNode struct {
	Next *MyListNode
	Val  int
}

type MyLinkedList struct {
	dummyHead *MyListNode
	size      int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &MyListNode{},
		size:      0,
	}
}

func (list *MyLinkedList) Get(index int) int {
	if index < 0 || index >= list.size {
		return -1
	}

	curNode := list.dummyHead.Next
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}

	return curNode.Val
}

func (list *MyLinkedList) AddAtHead(val int) {
	list.AddAtIndex(0, val)
}

func (list *MyLinkedList) AddAtTail(val int) {
	list.AddAtIndex(list.size, val)
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > list.size {
		return
	}

	curNode := list.dummyHead
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}

	originalNext := curNode.Next
	newNode := &MyListNode{Next: originalNext, Val: val}
	curNode.Next = newNode
	list.size++
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	// 防止溢出
	if index < 0 || index >= list.size {
		return
	}

	curNode := list.dummyHead
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}

	// 中间那个是需要删除的节点
	nextNode := curNode.Next.Next
	curNode.Next = nextNode
	list.size--
}

func main() {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	println(list.Get(1))
	list.DeleteAtIndex(1)
	println(list.Get(1))
}
