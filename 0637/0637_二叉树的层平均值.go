package averageOfLevelsInBinaryTree

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
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	l := []*TreeNode{root}

	for len(l) > 0 {
		size := len(l)
		sum := 0
		for i := 0; i < size; i++ {
			node := l[0]
			l = l[1:]
			sum += node.Val
			if node.Left != nil {
				l = append(l, node.Left)
			}

			if node.Right != nil {
				l = append(l, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(size))
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
