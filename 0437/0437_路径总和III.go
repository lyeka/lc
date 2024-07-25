package pathSumIii

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := 0
	ans += rootSum(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}

// rootSum 表示以节点 p 为起点向下且满足路径总和为
func rootSum(root *TreeNode, targetSum int) (ans int) {
	if root == nil {
		return
	}

	if root.Val == targetSum {
		ans++
	}
	ans += rootSum(root.Left, targetSum-root.Val)
	ans += rootSum(root.Right, targetSum-root.Val)

	return
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
