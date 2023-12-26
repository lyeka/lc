package validPalindrome

import (
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !is(s[left]) {
			left++
			continue
		}
		if !is(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func is(s byte) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9')

}

//leetcode submit region end(Prohibit modification and deletion)
