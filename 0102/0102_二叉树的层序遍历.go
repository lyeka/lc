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

	res := make([][]int, 0)
	l := []*TreeNode{root}

	for len(l) > 0 {
		valList := make([]int, 0)
		size := len(l)
		for size > 0 {
			node := l[0]
			l = l[1:]
			valList = append(valList, node.Val)
			if node.Left != nil {
				l = append(l, node.Left)
			}
			if node.Right != nil {
				l = append(l, node.Right)
			}
			size--
		}
		res = append(res, valList)

	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
