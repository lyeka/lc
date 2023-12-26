package reverseWordsInAString

import (
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	wordList := make([]string, 0)
	start := -1

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if start != -1 {
				wordList = append(wordList, s[start:i])
				start = -1
			}
			continue
		} else {
			if start == -1 {
				start = i
			}
		}
	}

	if start != -1 {
		wordList = append(wordList, s[start:])
	}

	for i := 0; i < len(wordList)/2; i++ {
		wordList[i], wordList[len(wordList)-1-i] = wordList[len(wordList)-1-i], wordList[i]
	}

	return strings.Join(wordList, " ")

}

//leetcode submit region end(Prohibit modification and deletion)
