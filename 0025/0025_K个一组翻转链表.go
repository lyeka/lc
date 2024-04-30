package reverseNodesInKGroup

type ListNode struct {
	al   int
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseKGroupV2(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{
		al:   0,
		Next: head,
	}

	a, b := head, head

	for {
		for i := 0; i < k; i++ {
			if b == nil {
				break
			}
			b = b.Next
		}
	}

	return dummyHead.Next
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
