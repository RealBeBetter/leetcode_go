package main

import "sort"

// 1. 两数之和
// https://leetcode.cn/problems/two-sum/description
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))

	for idx, num := range nums {
		if otherIdx, ok := numMap[target-num]; ok {
			return []int{idx, otherIdx}
		}

		numMap[num] = idx
	}

	return []int{}
}

// 49. 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/description
func groupAnagrams(strs []string) [][]string {
	// 按照字母序排序即可
	m := make(map[string][]string, len(strs))
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return (bytes[i] - '0') < (bytes[j] - '0')
		})

		s := string(bytes)
		if _, ok := m[s]; !ok {
			m[s] = make([]string, 0)
		}

		m[s] = append(m[s], str)
	}

	res := make([][]string, 0, len(m))
	for _, val := range m {
		res = append(res, val)
	}

	return res
}

// 49. 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/description
func groupAnagramsWithCnt(strs []string) [][]string {
	// 因为只有小写字母，所以可以使用长度为 26 的数组
	m := make(map[[26]int][]string, len(strs))
	for _, str := range strs {
		cnt := make([]int, 26)
		for _, c := range str {
			cnt[c-'a']++
		}

		m[[26]int(cnt)] = append(m[[26]int(cnt)], str)
	}

	res := make([][]string, 0, len(m))
	for _, val := range m {
		res = append(res, val)
	}

	return res
}

// 128. 最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool, len(nums))
	for _, num := range nums {
		numSet[num] = true
	}

	res := 1
	for i := range numSet {
		// 只计算连续序列的起点
		if numSet[i-1] {
			continue
		}

		streak, num := 1, i
		for numSet[num+1] {
			streak++
			num++
		}

		res = max(res, streak)
	}

	return res
}
