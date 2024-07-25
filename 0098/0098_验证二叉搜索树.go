package validateBinarySearchTree

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
func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(node *TreeNode, lower *int, upper *int) bool {
	if node == nil {
		return true
	}

	if lower != nil && node.Val <= *lower {
		return false
	}

	if upper != nil && node.Val >= *upper {
		return false
	}

	return helper(node.Left, lower, &node.Val) && helper(node.Right, &node.Val, upper)
}

//leetcode submit region end(Prohibit modification and deletion)
