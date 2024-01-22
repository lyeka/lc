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
	res := []int{}
	l := []*TreeNode{root}

	for len(l) > 0 {
		last := l[len(l)-1]
		res = append(res, last.Val)
		tmp := make([]*TreeNode, 0)
		for _, node := range l {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		l = tmp
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
