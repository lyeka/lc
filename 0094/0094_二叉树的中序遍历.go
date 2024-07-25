package binaryTreeInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
// 迭代
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		root = root.Right

	}

	return ans
}

// 递归
func inorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	ans = append(ans, inorderTraversalV2(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversalV2(root.Right)...)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
