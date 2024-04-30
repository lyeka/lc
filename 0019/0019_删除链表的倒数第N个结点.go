package removeNthNodeFromEndOfList

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	pre, fast := dummyHead, dummyHead
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		pre = pre.Next
		fast = fast.Next
	}

	if pre.Next != nil {
		pre.Next = pre.Next.Next
	} else {
		pre.Next = nil
	}

	return dummyHead.Next

}

//leetcode submit region end(Prohibit modification and deletion)
