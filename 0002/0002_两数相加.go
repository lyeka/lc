package addTwoNumbers

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	x, y, flag := 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		} else {
			x = 0
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}

		t := x + y + flag
		flag = t / 10
		p.Next = &ListNode{
			Val:  t % 10,
			Next: nil,
		}
		p = p.Next

	}

	if flag != 0 {
		p.Next = &ListNode{Val: flag}
	}

	return head.Next

}

//leetcode submit region end(Prohibit modification and deletion)
