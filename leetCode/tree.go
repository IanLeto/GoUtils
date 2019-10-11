package leetCode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 翻转二叉树
// 思路
func invertTree(root *TreeNode) *TreeNode {
	if root == nil { //(1) 考虑异常情况
		return root
	}
	if root.Left == nil && root.Right == nil { // (2)考虑退出情况
		return root
	}
	root.Right, root.Left = root.Left, root.Right // (3)只关注某一层干啥
	root.Right = invertTree(root.Right)           // (4)向下递归
	root.Left = invertTree(root.Left)

	return root // (5)输出结果

}
