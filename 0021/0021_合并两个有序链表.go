package mergeTwoSortedLists

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	for p1 != nil {
		p.Next = p1
		p1 = p1.Next
		p = p.Next
	}

	for p2 != nil {
		p.Next = p2
		p2 = p2.Next
		p = p.Next
	}
	return head.Next

}

//leetcode submit region end(Prohibit modification and deletion)
