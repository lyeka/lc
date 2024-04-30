package linkedListCycleIi

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
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for slow != p {
				slow = slow.Next
				p = p.Next
			}

			return p
		}
	}

	return nil

}

//leetcode submit region end(Prohibit modification and deletion)

func detectCycleV2(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})

	p := head
	for p != nil {
		if _, ok := m[p]; ok {
			return p
		}
		m[p] = struct{}{}
		p = p.Next

	}

	return nil

}
