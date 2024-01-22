package minimumAbsoluteDifferenceInBst

import (
	"math"
)

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
func getMinimumDifference(root *TreeNode) int {
	pre, ans := -1, math.MaxInt
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 {
			if node.Val-pre < ans {
				ans = node.Val - pre
			}
		}

		pre = node.Val

		dfs(node.Right)
	}
	dfs(root)
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
