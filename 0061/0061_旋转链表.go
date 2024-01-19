package rotateList

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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	tail := head
	cnt := 1
	for tail.Next != nil {
		tail = tail.Next
		cnt++
	}
	tail.Next = head

	mod := k % cnt

	p := head
	for i := 0; i < mod; i++ {
		p = p.Next
	}

	newHead := p.Next
	p.Next = nil

	return newHead

}

//leetcode submit region end(Prohibit modification and deletion)
