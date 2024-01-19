package maximumDepthOfBinaryTree

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

func maxDepth(root *TreeNode) int {
	return maxDepthV2(root)

}

// 递归，深度遍历
func maxDepthV1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

func maxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	p := []*TreeNode{root}
	for len(p) > 0 {
		size := len(p)
		for size > 0 {
			node := p[0]
			p = p[1:]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
			size--
		}
		res += 1
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

//leetcode submit region end(Prohibit modification and deletion)
