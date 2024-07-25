package mergeKSortedLists

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeList(lists, 0, len(lists)-1)
}

func mergeList(lists []*ListNode, l, r int) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if l == r {
		return lists[l]
	}
	mid := (l + r) / 2

	return merge(mergeList(lists, l, mid), mergeList(lists, mid+1, r))

}

func merge(l1, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}
