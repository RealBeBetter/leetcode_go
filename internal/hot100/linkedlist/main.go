package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 234. 回文链表
// https://leetcode.cn/problems/palindrome-linked-list/description
func isPalindrome(head *ListNode) bool {
	// 简单遍历方法实现
	vals := make([]int, 0)
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}

	for i, j := 0, len(vals)-1; i < len(vals)/2; i, j = i+1, j-1 {
		if vals[i] == vals[j] {
			continue
		}
		return false
	}

	return true
}

// 234. 回文链表
// https://leetcode.cn/problems/palindrome-linked-list/description
func isPalindromeWithRecursion(head *ListNode) bool {
	// 递归方法
	front := head
	var recursionCheck func(node *ListNode) bool
	recursionCheck = func(node *ListNode) bool {
		if node != nil {
			// 这一步会一直走到最后一个节点，仔细思考一下
			if !recursionCheck(node.Next) {
				return false
			}

			if node.Val != front.Val {
				return false
			}

			// 相等则继续遍历下一个节点
			front = front.Next
		}

		return true
	}

	return recursionCheck(head)
}

// 234. 回文链表
// https://leetcode.cn/problems/palindrome-linked-list/description
func isPalindromeWithPointer(head *ListNode) bool {
	// 找到链表的中间节点，并反转后半部分链表
	if head == nil {
		return true
	}

	fastNode, slowNode := head, head
	for fastNode.Next != nil && fastNode.Next.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}

	// 翻转 slowNode 后半部分节点，并保留 slowNode
	halfStartNode := reverseList(slowNode.Next)

	// halfStartNode 就是后半链表的开始节点
	for halfStartNode != nil {
		if head.Val != halfStartNode.Val {
			return false
		}

		head = head.Next
		halfStartNode = halfStartNode.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}
