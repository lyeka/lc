package reverseLinkedListIi

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head

	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	leftNode := pre.Next
	rightNode := leftNode

	for i := 0; i < right-left; i++ {
		rightNode = rightNode.Next
	}

	succ := rightNode.Next

	rightNode.Next = nil

	reverse(leftNode)

	pre.Next = rightNode
	leftNode.Next = succ

	return dummyNode.Next

}

func reverse(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

//leetcode submit region end(Prohibit modification and deletion)
