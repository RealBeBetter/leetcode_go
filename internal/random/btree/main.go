package btree

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

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

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	level := []*TreeNode{root}
	for len(level) > 0 {
		res++
		next := make([]*TreeNode, 0)
		for i := 0; i < len(level); i++ {
			if level[i].Left != nil {
				next = append(next, level[i].Left)
			}
			if level[i].Right != nil {
				next = append(next, level[i].Right)
			}
		}

		level = next
	}

	return res
}

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree
func maxDepthWithRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		res = max(res, depth)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	// 开始遍历时，最低层数是 1
	depth := 1
	order(root, depth)
	return res
}

// 111. 二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	level := []*TreeNode{root}
	for len(level) > 0 {
		end := false
		next := make([]*TreeNode, 0)
		for i := 0; i < len(level); i++ {
			// 左右子节点都为空，表示某条路径已经结束
			if level[i].Left == nil && level[i].Right == nil {
				end = true
				break
			}

			if level[i].Left != nil {
				next = append(next, level[i].Left)
			}
			if level[i].Right != nil {
				next = append(next, level[i].Right)
			}
		}

		res++
		level = next

		if end {
			break
		}
	}

	return res
}

// 111. 二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree
func minDepthWithRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := math.MaxInt
	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		// 左右子节点都为空，表示某条路径已经结束
		if root.Left == nil && root.Right == nil {
			res = min(res, depth)
		}

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, 1)
	return res
}

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/description/
func invertTreeWithPre(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreeWithPre(root.Left)
	invertTreeWithPre(root.Right)

	return root
}

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/description/
func invertTreeWithPost(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertTreeWithPost(root.Left)
	invertTreeWithPost(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/description/
func invertTreeWithMid(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// 因为中间替换过一次，所以后面遍历的时候需要切换一下节点
	invertTreeWithMid(root.Left)
	root.Left, root.Right = root.Right, root.Left
	invertTreeWithMid(root.Left)

	return root
}

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/description/
func invertTreeWithLevel(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	level := []*TreeNode{root}
	for len(level) > 0 {
		next := make([]*TreeNode, 0, len(level)*2)
		for i := 0; i < len(level); i++ {
			level[i].Left, level[i].Right = level[i].Right, level[i].Left
			if level[i].Left != nil {
				next = append(next, level[i].Left)
			}

			if level[i].Right != nil {
				next = append(next, level[i].Right)
			}
		}

		level = next
	}

	return root
}

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/description/
func invertTreeWithLevelRecursion(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var order func(root *TreeNode)
	order = func(root *TreeNode) {
		if root == nil {
			return
		}

		root.Left, root.Right = root.Right, root.Left
		order(root.Left)
		order(root.Right)
	}

	order(root)
	return root
}

// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/description/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	level := []*TreeNode{root}
	for len(level) > 0 {
		next := make([]*TreeNode, 0)
		for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
			if level[i] == nil && level[j] == nil {
				continue
			}

			if level[i] == nil || level[j] == nil {
				return false
			}

			if level[i] != nil && level[j] != nil {
				if level[i].Val != level[j].Val {
					return false
				}
			}
		}

		for i := 0; i < len(level); i++ {
			if level[i] != nil {
				next = append(next, level[i].Left)
				next = append(next, level[i].Right)
			}
		}

		level = next
	}

	return true
}

// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/description/
func isSymmetricWithRecursion(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var compare func(left *TreeNode, right *TreeNode) bool
	compare = func(left *TreeNode, right *TreeNode) bool {
		// 终止条件，当节点均为空，或者有不相等的
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}

		outside := compare(left.Left, right.Right)
		inside := compare(left.Right, right.Left)
		return inside && outside
	}

	return compare(root.Left, root.Right)
}

// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/description/
func isSymmetricWithStack(root *TreeNode) bool {
	if root == nil {
		return false
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root.Left)
	stack = append(stack, root.Right)
	for len(stack) > 0 {
		length := len(stack)
		right := stack[length-1]
		left := stack[length-2]
		stack = stack[:length-2]

		if left == nil && right == nil {
			continue
		} else if left == nil || right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}

		stack = append(stack, left.Left)
		stack = append(stack, right.Right)
		stack = append(stack, left.Right)
		stack = append(stack, right.Left)
	}

	return true
}

// 222. 完全二叉树的节点个数
// https://leetcode.cn/problems/count-complete-tree-nodes/description/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var count func(root *TreeNode) int
	count = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftCnt := count(root.Left)
		rightCnt := count(root.Right)
		return leftCnt + rightCnt + 1
	}

	return count(root)
}

// 222. 完全二叉树的节点个数
// https://leetcode.cn/problems/count-complete-tree-nodes/description/
func countNodesWithLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	level := []*TreeNode{root}
	for len(level) > 0 {
		next := make([]*TreeNode, 0)
		for i := 0; i < len(level); i++ {
			if level[i].Left != nil {
				next = append(next, level[i].Left)
			}
			if level[i].Right != nil {
				next = append(next, level[i].Right)
			}
			ans++
		}

		level = next
	}

	return ans
}

// 110. 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/description/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var calcHeight func(root *TreeNode) int
	calcHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftHeight := calcHeight(root.Left)
		rightHeight := calcHeight(root.Right)
		if leftHeight == -1 || rightHeight == -1 {
			return -1
		}

		// 相差大于 1，直接返回
		if math.Abs(float64(leftHeight-rightHeight)) > 1 {
			return -1
		}

		return max(leftHeight, rightHeight) + 1
	}

	return calcHeight(root) != -1
}

// 257. 二叉树的所有路径
// https://leetcode.cn/problems/binary-tree-paths
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := make([]string, 0)
	var order func(root *TreeNode, path []string, depth int)
	order = func(root *TreeNode, path []string, depth int) {
		if depth == len(path) {
			path = append(path, strconv.Itoa(root.Val))
		}

		if root.Left == nil && root.Right == nil {
			res = append(res, strings.Join(path, "->"))
			return
		}

		if root.Left != nil {
			order(root.Left, path, depth+1)
		}
		if root.Right != nil {
			order(root.Right, path, depth+1)
		}
	}

	order(root, []string{strconv.Itoa(root.Val)}, 0)
	return res
}

// 257. 二叉树的所有路径
// https://leetcode.cn/problems/binary-tree-paths
func binaryTreePathsWithBackTrack(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := make([]string, 0)
	var order func(root *TreeNode, path []string)
	order = func(root *TreeNode, path []string) {
		path = append(path, strconv.Itoa(root.Val))
		if root.Left == nil && root.Right == nil {
			res = append(res, strings.Join(path, "->"))
			return
		}

		if root.Left != nil {
			order(root.Left, path)
		}
		if root.Right != nil {
			order(root.Right, path)
		}
	}

	order(root, make([]string, 0))
	return res
}

// 100. 相同的树
// https://leetcode.cn/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	same := true

	var order func(p *TreeNode, q *TreeNode)
	order = func(p *TreeNode, q *TreeNode) {
		if p == nil && q == nil {
			return
		}

		if p == nil || q == nil || p.Val != q.Val {
			same = false
			return
		}

		order(p.Left, q.Left)
		order(p.Right, q.Right)
	}

	order(p, q)
	return same
}

// 100. 相同的树
// https://leetcode.cn/problems/same-tree/
func isSameTreeWithRecursion(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	var isSame func(p *TreeNode, q *TreeNode) bool
	isSame = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil || p.Val != q.Val {
			return false
		}

		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return isSame(p, q)
}

// 572. 另一棵树的子树
// https://leetcode.cn/problems/subtree-of-another-tree/description/
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	var isSame func(p *TreeNode, q *TreeNode) bool
	isSame = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil || p.Val != q.Val {
			return false
		}

		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	var isSub func(p *TreeNode, q *TreeNode) bool
	isSub = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil {
			return false
		}

		// 判断是否相同，是否为左子树，是否为右子树
		return isSame(p, q) || isSub(p.Left, q) || isSub(p.Right, q)
	}

	return isSub(root, subRoot)
}

// 404. 左叶子之和
// https://leetcode.cn/problems/sum-of-left-leaves/description/
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var order func(root *TreeNode) int
	order = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		if root.Left == nil && root.Right == nil {
			return 0
		}

		// 左侧符合左叶子条件，则终止左侧递归
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			return root.Left.Val + order(root.Right)
		}

		return order(root.Left) + order(root.Right)
	}

	return order(root)
}

// 404. 左叶子之和
// https://leetcode.cn/problems/sum-of-left-leaves/description/
func sumOfLeftLeavesWithLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0

	level := []*TreeNode{root}

	for len(level) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for i := 0; i < len(level); i++ {
			node := level[i]
			if node.Left != nil {
				if node.Left.Left == nil && node.Left.Right == nil {
					sum += node.Left.Val
				} else {
					nextLevel = append(nextLevel, node.Left)
				}
			}

			if node.Right != nil {
				if !(node.Right.Left == nil && node.Right.Right == nil) {
					nextLevel = append(nextLevel, node.Right)
				}
			}
		}

		level = nextLevel
	}

	return sum
}

// 513. 找树左下角的值
// https://leetcode.cn/problems/find-bottom-left-tree-value
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0

	level := []*TreeNode{root}
	for len(level) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for i := 0; i < len(level); i++ {
			if i == 0 && level[i] != nil {
				ans = level[i].Val
			}

			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}

			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}

		level = nextLevel
	}

	return ans
}

// 513. 找树左下角的值
// https://leetcode.cn/problems/find-bottom-left-tree-value
func findBottomLeftValueWithRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}

	globalDepth, ans := 0, 0

	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if globalDepth == depth {
			globalDepth++
			ans = root.Val
		}

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, 0)
	return ans
}

// 112. 路径总和
// https://leetcode.cn/problems/path-sum
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	res := false

	// 使用中序遍历
	var order func(root *TreeNode, sum int)
	order = func(root *TreeNode, sum int) {
		sum += root.Val
		if root.Left == nil && root.Right == nil {
			if sum == targetSum {
				res = true
				return
			}
		}

		if root.Left != nil {
			order(root.Left, sum)
		}

		if root.Right != nil {
			order(root.Right, sum)
		}
	}

	order(root, 0)
	return res
}

// 112. 路径总和
// https://leetcode.cn/problems/path-sum
func hasPathSumWithBackTrace(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var traverse func(root *TreeNode, sum int) bool
	traverse = func(root *TreeNode, sum int) bool {
		if root.Left == nil && root.Right == nil {
			return sum == 0
		}

		if root.Left != nil {
			sum -= root.Left.Val
			if traverse(root.Left, sum) {
				return true
			}
			sum += root.Left.Val
		}

		if root.Right != nil {
			sum -= root.Right.Val
			if traverse(root.Right, sum) {
				return true
			}
			sum += root.Right.Val
		}

		return false
	}

	return traverse(root, targetSum-root.Val)
}

// 113. 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/description/
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	var order func(root *TreeNode, path *[]int, sum int)
	order = func(root *TreeNode, path *[]int, sum int) {
		sum += root.Val
		*path = append(*path, root.Val)

		if root.Left == nil && root.Right == nil {
			if sum == targetSum {
				temp := make([]int, 0, len(*path))
				for _, val := range *path {
					temp = append(temp, val)
				}
				ans = append(ans, temp)
			}
			return
		}

		if root.Left != nil {
			order(root.Left, path, sum)
			*path = (*path)[:len(*path)-1]
		}

		if root.Right != nil {
			order(root.Right, path, sum)
			*path = (*path)[:len(*path)-1]
		}
	}

	path := make([]int, 0)
	order(root, &path, 0)
	return ans
}
