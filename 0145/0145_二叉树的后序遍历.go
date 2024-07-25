package binaryTreePostorderTraversal

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == pre {
			ans = append(ans, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}

	}
	return ans
}

// 递归
func postorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	ans = append(ans, postorderTraversalV2(root.Left)...)
	ans = append(ans, postorderTraversalV2(root.Right)...)
	ans = append(ans, root.Val)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
