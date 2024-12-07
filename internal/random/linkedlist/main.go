package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 203. 移除链表元素
// https://leetcode.cn/problems/remove-linked-list-elements
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val != val {
			cur = cur.Next
			continue
		}

		cur.Next = cur.Next.Next
	}

	return dummyHead.Next
}

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	node := head
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	return prev
}

func reverseListWithRecursion(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}

	next := cur.Next
	cur.Next = pre
	return reverse(cur, next)
}

// 24. 两两交换链表中的节点
// https://leetcode.cn/problems/swap-nodes-in-pairs
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	// 引入虚拟头结点，让所有问题一致化
	curNode := dummyHead
	// curNode -> 1 -> 2 -> 3
	// 要将其中的 1 与 2 交换，3 是 nextNode
	for curNode.Next != nil && curNode.Next.Next != nil {
		firstNode := curNode.Next
		secondNode := curNode.Next.Next
		// 记录下一轮循环中的起始节点，不能在这里断开
		nextNode := secondNode.Next

		curNode.Next = secondNode
		curNode.Next.Next = firstNode
		curNode = curNode.Next.Next
		curNode.Next = nextNode
	}

	return dummyHead.Next
}

// 19. 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	fastNode, slowNode := dummyHead, dummyHead
	for i := 0; i < n; i++ {
		fastNode = fastNode.Next
	}

	// 这时候 fastNode 到尽头时，slowNode 刚好是要删除的节点
	// 因此 fastNode 还要进一步
	fastNode = fastNode.Next

	for fastNode != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}

	slowNode.Next = slowNode.Next.Next
	return dummyHead.Next
}

// 160. 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 将尾端对齐，判断第一个相等的节点
	lenA, lenB := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}

	// 计算长度之后，还原节点位置
	curA, curB = headA, headB

	// 移动更长的链表指针，使其与短的链表保持一致的位置
	lenDiff := int(math.Abs(float64(lenA - lenB)))
	for lenDiff > 0 {
		if lenA > lenB {
			curA = curA.Next
		} else {
			curB = curB.Next
		}

		lenDiff--
	}

	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}

		curA = curA.Next
		curB = curB.Next
	}

	return nil
}

// 160. 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists
func getIntersectionNodeWithSet(headA, headB *ListNode) *ListNode {
	// 因为 node 指针可比较，所以可以利用 set 来寻址
	set := make(map[*ListNode]struct{})
	curA, curB := headA, headB

	for curA != nil {
		set[curA] = struct{}{}
		curA = curA.Next
	}

	for curB != nil {
		_, ok := set[curB]
		if !ok {
			curB = curB.Next
			continue
		}

		return curB
	}

	return nil
}

// 160. 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists
// 参考解法：https://leetcode.cn/problems/intersection-of-two-linked-lists/solutions/793580/xiang-jiao-lian-biao-hashmap-shuang-zhi-9w3uh/
func getIntersectionNodeWithDoublePtr(headA, headB *ListNode) *ListNode {
	// 初始化指针
	curA, curB := headA, headB

	// 在两次遍历中，切换链表的操作，会使得两者走过的路都是一样的
	// 如果有交叉点，则会在交叉点相遇
	// 如果没有交叉点，则最后都会在为 nil 的时候跳出循环
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}

	return curA
}

// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii
func detectCycle(head *ListNode) *ListNode {
	slowNode, fastNode := head, head
	for fastNode != nil && fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
		// 如果有环，那么快慢指针始终会相遇
		if slowNode == fastNode {
			// 将其中一个节点移至头节点
			slowNode = head
			for slowNode != fastNode {
				fastNode = fastNode.Next
				slowNode = slowNode.Next
			}

			// 二次相遇的节点，即是头节点
			return slowNode
		}
	}

	return nil
}

// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii
func detectCycleWithSet(head *ListNode) *ListNode {
	set := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := set[head]; ok {
			return head
		}

		set[head] = struct{}{}
		head = head.Next
	}

	return nil
}
