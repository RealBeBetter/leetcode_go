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
