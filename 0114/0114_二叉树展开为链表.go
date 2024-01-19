package flattenBinaryTreeToLinkedList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	right := root.Right
	root.Right = root.Left
	root.Left = nil
	p := root
	for p.Right != nil {
		p = p.Right
	}

	p.Right = right
}

//leetcode submit region end(Prohibit modification and deletion)
