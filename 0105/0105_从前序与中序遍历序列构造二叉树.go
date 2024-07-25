package constructBinaryTreeFromPreorderAndInorderTraversal

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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	index := 0
	for i, v := range inorder {
		if v == rootVal {
			index = i
			break
		}
	}

	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(preorder[1:1+index], inorder[:index])
	root.Right = buildTree(preorder[1+index:], inorder[index+1:])
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
