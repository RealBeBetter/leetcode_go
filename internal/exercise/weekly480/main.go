package weekly480

import (
	"sort"
	"strings"
)

// Q1. 最大和最小 K 个元素的绝对差
// https://leetcode.cn/contest/weekly-contest-480/problems/absolute-difference-between-maximum-and-minimum-k-elements/
func absDifference(nums []int, k int) int {
	sort.Ints(nums)
	minSum, maxSum := 0, 0
	for i, j := 0, len(nums)-1; i < k; i, j = i+1, j-1 {
		minSum += nums[i]
		maxSum += nums[j]
	}

	if maxSum-minSum > 0 {
		return maxSum - minSum
	}
	return minSum - maxSum
}

// Q2. 反转元音数相同的单词
// https://leetcode.cn/contest/weekly-contest-480/problems/reverse-words-with-same-vowel-count/
func reverseWords(s string) string {
	isVowel := func(c byte) bool {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			return true
		}

		return false
	}

	calcVowelCnt := func(word string) int {
		vowelCnt := 0
		for _, c := range []byte(word) {
			if isVowel(c) {
				vowelCnt++
			}
		}
		return vowelCnt
	}

	words := strings.Split(s, " ")
	vowelCnt := calcVowelCnt(words[0])

	ans := strings.Builder{}
	ans.WriteString(words[0])
	for i := 1; i < len(words); i++ {
		cnt := calcVowelCnt(words[i])
		if cnt == vowelCnt {
			ans.WriteString(" ")
			for j := len(words[i]) - 1; j >= 0; j-- {
				ans.WriteByte(words[i][j])
			}
		} else {
			ans.WriteString(" " + words[i])
		}
	}

	return ans.String()
}

// Q3. 使循环数组余额非负的最少移动次数
// https://leetcode.cn/contest/weekly-contest-480/problems/minimum-moves-to-balance-circular-array/
func minMoves(balance []int) int64 {
	sum := 0
	for _, num := range balance {
		sum += num
	}
	if sum < 0 {
		return -1
	}

	moveBalance := func(origin, target, distance int) int {
		if balance[origin] <= 0 {
			return 0
		}

		if balance[origin]+balance[target] >= 0 {
			cost := -balance[target] * distance
			balance[origin] += balance[target]
			balance[target] = 0
			return cost
		}

		cost := balance[origin] * distance
		balance[target] += balance[origin]
		balance[origin] = 0
		return cost
	}

	calcIdx := func(idx int) int {
		if idx < 0 {
			return idx + len(balance)
		} else if idx > len(balance)-1 {
			return idx - len(balance)
		}
		return idx
	}

	ans := 0
	// 填平一个负数，需要的总次数是：(distance * num) * n
	for i := 0; i < len(balance); i++ {
		if balance[i] >= 0 {
			continue
		}

		// 从左或者从右开始遍历，保持左右两侧距离相等，同距离优先取左边，不能取已经为 0 或者不存在的数
		for distance := 1; balance[i] < 0; distance++ {
			// 计算左右 idx 值
			leftIdx := calcIdx(i - distance)
			rightIdx := calcIdx(i + distance)

			// 先取左边
			cost := moveBalance(leftIdx, i, distance)
			ans += cost
			if balance[i] == 0 {
				break
			}

			// 再取右边
			cost = moveBalance(rightIdx, i, distance)
			ans += cost
		}
	}

	return int64(ans)
}
