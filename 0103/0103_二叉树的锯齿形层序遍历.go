package binaryTreeZigzagLevelOrderTraversal

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	l := []*TreeNode{root}

	fromRight := false

	for len(l) > 0 {
		vals := make([]int, 0)
		size := len(l)
		for i := 0; i < size; i++ {
			valIndex := i
			if fromRight {
				valIndex = size - 1 - i
			}
			if l[i].Left != nil {
				l = append(l, l[i].Left)
			}
			if l[i].Right != nil {
				l = append(l, l[i].Right)
			}
			vals = append(vals, l[valIndex].Val)
		}

		fromRight = !fromRight
		l = l[size:]
		res = append(res, vals)

	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
