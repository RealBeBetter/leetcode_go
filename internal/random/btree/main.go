package btree

import "slices"

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 0)
	preTraverse(&res, root)
	return res
}

func preTraverse(res *[]int, node *TreeNode) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)
	preTraverse(res, node.Left)
	preTraverse(res, node.Right)
}

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/
func preorderTraversalByIterate(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/
func preorderTraversalByUnified(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node != nil {
			// 将该节点弹出，避免重复操作
			stack = stack[:len(stack)-1]

			// 反向入栈
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			// 根节点暂未处理过，需要入栈标记
			stack = append(stack, node)
			stack = append(stack, nil)
			continue
		}

		// 遇到空节点，则空节点与后面的节点需要处理
		stack = stack[:len(stack)-1]
		node = stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
	}

	return res
}

// 94. 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 0)
	inTraverse(&res, root)
	return res
}

func inTraverse(res *[]int, node *TreeNode) {
	if node == nil {
		return
	}

	inTraverse(res, node.Left)
	*res = append(*res, node.Val)
	inTraverse(res, node.Right)
}

// 94. 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal
func inorderTraversalByIterate(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
			continue
		}

		// node 为空时，需要节点出栈
		node = stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		node = node.Right
	}

	return res
}

// 94. 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal
func inorderTraversalByUnified(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node != nil {
			// 将该节点弹出，避免重复操作
			stack = stack[:len(stack)-1]

			// 反向入栈
			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			// 根节点暂未处理过，需要入栈标记
			stack = append(stack, node)
			stack = append(stack, nil)

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			continue
		}

		// 遇到空节点，则空节点与后面的节点需要处理
		stack = stack[:len(stack)-1]
		node = stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
	}

	return res
}

// 145. 二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postTraverse(&res, root)
	return res
}

func postTraverse(res *[]int, node *TreeNode) {
	if node == nil {
		return
	}

	postTraverse(res, node.Left)
	postTraverse(res, node.Right)
	*res = append(*res, node.Val)
}

// 145. 二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal/
func postorderTraversalByIterate(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	// 这时候是根右左，需要反转
	slices.Reverse(res)
	return res
}

// 145. 二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal/
func postorderTraversalByUnified(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node != nil {
			// 将该节点弹出，避免重复操作
			stack = stack[:len(stack)-1]

			// 根节点暂未处理过，需要入栈标记
			stack = append(stack, node)
			stack = append(stack, nil)

			// 反向入栈
			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			continue
		}

		// 此时为标记的空节点，需要出栈
		stack = stack[:len(stack)-1]
		node = stack[len(stack)-1]
		res = append(res, node.Val)
		// 已处理过，需要出栈
		stack = stack[:len(stack)-1]
	}

	return res
}

// 102. 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	curLevel := []*TreeNode{root}

	for len(curLevel) > 0 {
		curValues := make([]int, 0)
		nextLevel := make([]*TreeNode, 0)
		for i := 0; i < len(curLevel); i++ {
			if curLevel[i] == nil {
				continue
			}

			curValues = append(curValues, curLevel[i].Val)
			nextLevel = append(nextLevel, curLevel[i].Left)
			nextLevel = append(nextLevel, curLevel[i].Right)
		}

		curLevel = nextLevel
		if len(curValues) > 0 {
			res = append(res, curValues)
		}
	}

	return res
}

// 102. 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/
func levelOrderWithRecursion(root *TreeNode) [][]int {
	var res [][]int

	depth := 0

	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(res) == depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	return res
}
