package binaryTreeLevelOrderTraversal

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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := make([][]int, 0)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		size := len(stack)
		vals := make([]int, 0)
		for size > 0 {
			p := stack[0]
			stack = stack[1:]
			vals = append(vals, p.Val)
			if p.Left != nil {
				stack = append(stack, p.Left)
			}
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			size--
		}
		ans = append(ans, vals)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
