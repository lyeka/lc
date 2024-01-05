package minimumWindowSubstring

//leetcode submit region begin(Prohibit modification and deletion)
// 小技巧
// 1. 利用 [128]int{} 替代 map
// 2. 利用 s-'A' 来进一步减少 []int 大小，128->58
func minWindow(s string, t string) string {
	//ori, cnt := map[byte]int{}, map[byte]int{}
	ori, cnt := [58]int{}, [58]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]-'A']++
	}

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	length := len(s) + 1
	ansL, ansR := -1, -1

	for l, r := 0, 0; r < len(s); r++ {
		if ori[s[r]-'A'] > 0 {
			cnt[s[r]-'A']++
		}
		for check() && l <= r {
			if (r - l + 1) < length {
				length = r - l + 1
				ansL, ansR = l, r
			}
			if cnt[s[l]-'A'] > 0 {
				cnt[s[l]-'A']--
			}

			l++
		}
	}
	if ansL == -1 {
		return ""
	}

	return s[ansL : ansR+1]

}

//leetcode submit region end(Prohibit modification and deletion)

func minWindowV2(s string, t string) string {
	cnt := [58]int{}
	need := [58]int{}
	needCnt := 0
	for i := range t {
		if need[t[i]-'A'] == 0 {
			needCnt++
		}
		need[t[i]-'A']++
	}

	l, minLen := 0, len(s)+1
	// 存储最小子串的开始下标
	res := 0
	for r := range s {
		// 更新下标 right
		if need[s[r]-'A'] == 0 {
			continue
		}
		cnt[s[r]-'A']++
		if cnt[s[r]-'A'] == need[s[r]-'A'] {
			needCnt--
		}
		// 匹配到所有字符
		for needCnt == 0 {
			if need[s[l]-'A'] == 0 {
				l++
				continue
			}
			// 只处理 t 中存在的字符
			cnt[s[l]-'A']--
			if cnt[s[l]-'A'] < need[s[l]-'A'] {
				needCnt++
				if r-l+1 < minLen {
					minLen = r - l + 1
					res = l
				}
			}
			l++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[res : res+minLen]
}
