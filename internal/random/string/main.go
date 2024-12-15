package main

import (
	"strings"
)

// 344. 反转字符串
// https://leetcode.cn/problems/reverse-string
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 541. 反转字符串 II
// https://leetcode.cn/problems/reverse-string-ii
func reverseStr(s string, k int) string {
	bytes := []byte(s)

	for i := 0; i < len(s); i += 2 * k {
		remainingLen := len(s) - i
		if remainingLen < k {
			reverseString(bytes[i:])
			break
		}

		reverseString(bytes[i : i+k])

		if remainingLen < 2*k {
			break
		}
	}

	return string(bytes)
}

// 648. 单词替换
// https://leetcode.cn/problems/replace-words
func replaceWords(dictionary []string, sentence string) string {
	str := strings.Builder{}
	str.Grow(len(sentence))

	words := strings.Split(sentence, " ")
	for i, word := range words {
		prefix := word
		for _, dict := range dictionary {
			if !strings.HasPrefix(word, dict) {
				continue
			}

			if len(dict) < len(prefix) {
				prefix = dict
			}
		}

		str.Write([]byte(prefix))
		if i != len(words)-1 {
			str.Write([]byte(" "))
		}
	}

	return str.String()
}

// 648. 单词替换
// https://leetcode.cn/problems/replace-words
func replaceWordsWithTrie(dictionary []string, sentence string) string {
	type trie map[rune]trie
	root := trie{}
	// 构建 trie 树
	for _, dict := range dictionary {
		cur := root
		for _, d := range dict {
			if cur[d] == nil {
				cur[d] = trie{}
			}

			cur = cur[d]
		}

		cur['#'] = trie{}
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, w := range word {
			if cur['#'] != nil {
				// 已找到最短的前缀
				words[i] = word[:j]
				break
			}

			if cur[w] == nil {
				// 未找到匹配的前缀
				break
			}

			cur = cur[w]
		}
	}

	return strings.Join(words, " ")
}

// 151. 反转字符串中的单词
// https://leetcode.cn/problems/reverse-words-in-a-string
func reverseWords(s string) string {
	wordsWithBlank := strings.Split(s, " ")
	words := make([]string, 0, len(wordsWithBlank))
	for i := 0; i < len(wordsWithBlank); i++ {
		if strings.TrimSpace(wordsWithBlank[i]) == "" {
			continue
		}
		words = append(words, wordsWithBlank[i])
	}

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

// 28. 找出字符串中第一个匹配项的下标
// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}

	// 使用 KMP 算法实现
	next := calcNext(needle)
	// 模式串起始的下标位置
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			// 回退到 next 数组的前一位
			j = next[j-1]
		}
		// 每次前移的只有模式串
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	return -1
}

func calcNext(s string) []int {
	next := make([]int, len(s))
	// i 为后缀，j 为前缀
	for i := 1; i < len(s); i++ {
		prefix := []byte(s[:i])
		suffix := []byte(s[1 : i+1])
		next[i] = compare(prefix, suffix)
	}

	return next
}

func compare(a, b []byte) int {
	// 比对前后缀是否相等，算出相等先后缀的长度
	for len(a) > 0 && len(b) > 0 {
		sameCnt := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				// 不相等则再取各自的前后缀
				a = a[:len(a)-1]
				b = b[1:]
				sameCnt = 0
				break
			}

			sameCnt++
		}

		if sameCnt == len(a) {
			return len(a)
		}
	}

	return 0
}

// 28. 找出字符串中第一个匹配项的下标
// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string
func strStrWithNext(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}

	// 使用 KMP 算法实现
	next := calcNextArr(needle)
	// 模式串起始的下标位置
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			// 回退到 next 数组的前一位
			j = next[j-1]
		}
		// 每次前移的只有模式串
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	return -1
}

func calcNextArr(needle string) []int {
	next := make([]int, len(needle))
	j := 0
	for i := 1; i < len(next); i++ {
		for j > 0 && needle[i] != needle[j] {
			// 回退到前一位
			j = next[j-1]
		}

		if needle[i] == needle[j] {
			j++
		}

		// 当前最长相等前后缀长度
		next[i] = j
	}

	return next
}

// 459. 重复的子字符串
// https://leetcode.cn/problems/repeated-substring-pattern/
func repeatedSubstringPattern(s string) bool {
	// 使用查询方法
	if len(s) <= 1 {
		return false
	}

	repeatStr := s[1:] + s[:len(s)-1]
	return strings.Contains(repeatStr, s)
}
