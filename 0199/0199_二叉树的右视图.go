package binaryTreeRightSideView

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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	ans := make([]int, 0)
	for len(stack) > 0 {
		ans = append(ans, stack[len(stack)-1].Val)
		tmp := make([]*TreeNode, 0)
		for _, node := range stack {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		stack = tmp
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
