package copyListWithRandomPointer

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var m map[*Node]*Node

func deepCopy(n *Node) *Node {
	if n == nil {
		return nil
	}
	newNode := &Node{
		Val: n.Val,
	}
	m[n] = newNode
	newNode.Next = deepCopy(n.Next)
	newNode.Random = m[n.Random]
	return newNode
}

func copyRandomList(head *Node) *Node {
	m = make(map[*Node]*Node)
	return deepCopy(head)
}

//leetcode submit region end(Prohibit modification and deletion)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
