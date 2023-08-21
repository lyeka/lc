package removeNthNodeFromEndOfList

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	s, f := head, head
	for i := 0; i < n; i++ {
		f = f.Next
	}

	for f != nil && f.Next != nil {
		f = f.Next
		s = s.Next
	}

	if f == nil {
		return s.Next
	}

	next := s.Next.Next
	s.Next = next

	return head

}

//leetcode submit region end(Prohibit modification and deletion)
