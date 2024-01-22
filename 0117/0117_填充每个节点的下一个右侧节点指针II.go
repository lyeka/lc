package populatingNextRightPointersInEachNodeIi

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	list := []*Node{root}
	for len(list) > 0 {
		size := len(list)
		for size > 0 {
			node := list[0]
			list = list[1:]

			if size > 1 {
				node.Next = list[0]
			}

			if node.Left != nil {
				list = append(list, node.Left)
			}

			if node.Right != nil {
				list = append(list, node.Right)
			}

			size--
		}
	}

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
