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
	head := &ListNode{}
	p := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = &ListNode{
				Val:  list1.Val,
				Next: nil,
			}
			list1 = list1.Next
		} else {
			p.Next = &ListNode{
				Val:  list2.Val,
				Next: nil,
			}
			list2 = list2.Next
		}
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	}

	if list2 != nil {
		p.Next = list2
	}

	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
