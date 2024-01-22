package sumRootToLeafNumbers

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

func dfs(root *TreeNode, preNum int) int {
	if root == nil {
		return 0
	}

	sum := preNum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return dfs(root.Left, sum) + dfs(root.Right, sum)

}
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
