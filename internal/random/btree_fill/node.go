package btree_fill

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116. 填充每个节点的下一个右侧节点指针
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	res := root
	level := []*Node{root}
	for len(level) > 0 {
		// 规定输入是完美二叉树
		nextLevel := make([]*Node, 0)
		curNode := level[0]
		for i := 0; i < len(level); i++ {
			if i > 0 {
				curNode.Next = level[i]
				curNode = level[i]
			}

			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}
			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}

		// 输入中 Next 指针默认为 nil，这一步也可以省略
		curNode.Next = nil
		level = nextLevel
	}

	return res
}

// 116. 填充每个节点的下一个右侧节点指针
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node
func connectWithRecursion(root *Node) *Node {
	if root == nil {
		return root
	}

	res := root
	nodes := make([][]*Node, 0)

	var order func(root *Node, depth int)
	order = func(root *Node, depth int) {
		if root == nil {
			return
		}

		if len(nodes) == depth {
			// 新一层的起点，初始化这一层的变量
			nodes = append(nodes, []*Node{})
			// 上一层的结尾节点的 Next 默认是 Nil，可以不用处理；下一步可以省略
			if depth > 0 {
				nodes[depth-1][len(nodes[depth-1])-1].Next = nil
			}
		} else {
			lastIdx := len(nodes[depth]) - 1
			nodes[depth][lastIdx].Next = root
		}

		// 添加该节点
		nodes[depth] = append(nodes[depth], root)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	depth := 0
	order(root, depth)

	return res
}

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii
func connectII(root *Node) *Node {
	if root == nil {
		return root
	}

	res := root
	level := []*Node{root}
	for len(level) > 0 {
		nextLevel := make([]*Node, 0)
		curNode := level[0]
		for i := 0; i < len(level); i++ {
			if i > 0 {
				curNode.Next = level[i]
				curNode = level[i]
			}

			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}
			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}

		// 输入中 Next 指针默认为 nil，这一步也可以省略
		curNode.Next = nil
		level = nextLevel
	}

	return res
}

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii
func connectIIWithRecursion(root *Node) *Node {
	if root == nil {
		return root
	}

	res := root
	nodes := make([][]*Node, 0)

	var order func(root *Node, depth int)
	order = func(root *Node, depth int) {
		if root == nil {
			return
		}

		if len(nodes) == depth {
			// 新一层的起点，初始化这一层的变量
			nodes = append(nodes, []*Node{})
			// 上一层的结尾节点的 Next 默认是 Nil，可以不用处理；下一步可以省略
			if depth > 0 {
				nodes[depth-1][len(nodes[depth-1])-1].Next = nil
			}
		} else {
			lastIdx := len(nodes[depth]) - 1
			nodes[depth][lastIdx].Next = root
		}

		// 添加该节点
		nodes[depth] = append(nodes[depth], root)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	depth := 0
	order(root, depth)

	return res
}
