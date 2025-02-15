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

// 107. 二叉树的层序遍历 II
// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	level := []*TreeNode{root}

	for len(level) > 0 {
		next := make([]*TreeNode, 0, 2*len(level))
		values := make([]int, 0, 2*len(level))
		for i := 0; i < len(level); i++ {
			if level[i] == nil {
				continue
			}

			values = append(values, level[i].Val)
			next = append(next, level[i].Left)
			next = append(next, level[i].Right)
		}

		if len(values) > 0 {
			res = append(res, values)
		}

		level = next
	}

	slices.Reverse(res)
	return res
}

// 107. 二叉树的层序遍历 II
// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottomWithRecursion(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	depth := 0
	res := make([][]int, 0)

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
	slices.Reverse(res)

	return res
}

// 199. 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	level := []*TreeNode{root}
	for len(level) > 0 {
		var right *TreeNode

		next := make([]*TreeNode, 0, 2*len(level))
		for i := 0; i < len(level); i++ {
			if level[i] == nil {
				continue
			}

			right = level[i]
			next = append(next, level[i].Left)
			next = append(next, level[i].Right)
		}

		level = next
		if right != nil {
			res = append(res, right.Val)
		}
	}

	return res
}

// 199. 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideViewWithRecursion(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	depth := 0
	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(res) == depth {
			res = append(res, root.Val)
		} else {
			// 如果非第一个值，则表示非这一层的值
			// 那么后面遍历到的值，相比这个位置已有的值，更靠右边
			res[depth] = root.Val
		}

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	return res
}

// 637. 二叉树的层平均值
// https://leetcode.cn/problems/average-of-levels-in-binary-tree/
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	res := make([]float64, 0)
	level := []*TreeNode{root}
	for len(level) > 0 {
		next := make([]*TreeNode, 0, 2*len(level))
		sum := 0.0
		for i := 0; i < len(level); i++ {
			sum += float64(level[i].Val)
			if level[i].Left != nil {
				next = append(next, level[i].Left)
			}
			if level[i].Right != nil {
				next = append(next, level[i].Right)
			}
		}

		avg := sum / float64(len(level))
		res = append(res, avg)

		level = next
	}

	return res
}

// 637. 二叉树的层平均值
// https://leetcode.cn/problems/average-of-levels-in-binary-tree/
func averageOfLevelsWithRecursion(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	nums := make([][]float64, 0)

	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if depth == len(nums) {
			nums = append(nums, []float64{float64(root.Val)})
		} else {
			nums[depth] = append(nums[depth], float64(root.Val))
		}

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	depth := 0
	order(root, depth)

	res := make([]float64, len(nums))
	for i, num := range nums {
		sum := 0.0
		for _, n := range num {
			sum += n
		}

		res[i] = sum / float64(len(num))
	}

	return res
}

// 429. N 叉树的层序遍历
// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
func levelOrderOfNary(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	level := []*Node{root}
	for len(level) > 0 {
		nums := make([]int, 0, 2*len(level))
		next := make([]*Node, 0, 2*len(level))
		for i := 0; i < len(level); i++ {
			if level[i] == nil {
				continue
			}

			nums = append(nums, level[i].Val)
			next = append(next, level[i].Children...)
		}

		if len(nums) > 0 {
			res = append(res, nums)
		}
		level = next
	}

	return res
}

// 429. N 叉树的层序遍历
// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
func levelOrderOfNaryWithRecursion(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	depth := 0
	res := make([][]int, 0)

	var order func(root *Node, depth int)
	order = func(root *Node, depth int) {
		if root == nil {
			return
		}

		if len(res) == depth {
			res = append(res, []int{root.Val})
		} else {
			res[depth] = append(res[depth], root.Val)
		}

		children := root.Children
		for i := 0; i < len(children); i++ {
			order(children[i], depth+1)
		}
	}

	order(root, depth)

	return res
}

// 515. 在每个树行中找最大值
// https://leetcode.cn/problems/find-largest-value-in-each-tree-row
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	level := []*TreeNode{root}
	for len(level) > 0 {
		next := make([]*TreeNode, 0)

		maxVal := level[0].Val
		for i := 0; i < len(level); i++ {
			if level[i].Val > maxVal {
				maxVal = level[i].Val
			}

			if level[i].Left != nil {
				next = append(next, level[i].Left)
			}

			if level[i].Right != nil {
				next = append(next, level[i].Right)
			}
		}

		level = next
		res = append(res, maxVal)
	}

	return res
}

// 515. 在每个树行中找最大值
// https://leetcode.cn/problems/find-largest-value-in-each-tree-row
func largestValuesWithRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if depth == len(res) {
			res = append(res, root.Val)
		} else {
			if root.Val > res[depth] {
				res[depth] = root.Val
			}
		}

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	depth := 0
	order(root, depth)

	return res
}
