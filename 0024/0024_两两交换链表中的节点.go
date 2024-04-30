package swapNodesInPairs

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	return swapHelper(head, 2)

}

func swapHelper(head *ListNode, k int) *ListNode {
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = swapHelper(b, k)
	return newHead

}

func reverse(head, tail *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

//leetcode submit region end(Prohibit modification and deletion)

func swapPairsV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	start := dummyHead.Next
	p := dummyHead
	for start != nil && start.Next != nil {
		end := start.Next
		swap(start, end)
		p.Next = end
		p = start
		start = start.Next
	}

	return dummyHead.Next

}

func swap(s, e *ListNode) {
	if s == nil || e == nil {
		return
	}
	s.Next = e.Next
	e.Next = s

}

type ListNode struct {
	Val  int
	Next *ListNode
}
