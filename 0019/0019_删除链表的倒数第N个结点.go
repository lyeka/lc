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
	slow, fast := dummyHead, dummyHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	pre := slow
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}

	pre.Next = slow.Next
	return dummyHead.Next

}

//leetcode submit region end(Prohibit modification and deletion)
