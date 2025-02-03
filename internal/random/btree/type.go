package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node 多叉树节点
type Node struct {
	Val      int
	Children []*Node
}
