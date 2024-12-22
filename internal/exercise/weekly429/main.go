package main

import (
	"sort"
)

// Q1. 使数组元素互不相同所需的最少操作次数
// https://leetcode.cn/contest/weekly-contest-429/problems/minimum-number-of-operations-to-make-elements-in-array-distinct
func minimumOperations(nums []int) int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = m[nums[i]] + 1
	}

	ans := 0
	for hasSame(m) {
		if len(nums) < 3 {
			ans++
			break
		}

		for i := 0; i < 3; i++ {
			if m[nums[i]] == 1 {
				delete(m, nums[i])
			} else {
				m[nums[i]]--
			}
		}

		nums = nums[3:]
		ans++
	}

	return ans
}

func hasSame(m map[int]int) bool {
	for _, val := range m {
		if val > 1 {
			return true
		}
	}

	return false
}

// Q2. 执行操作后不同元素的最大数量
// https://leetcode.cn/contest/weekly-contest-429/problems/maximum-number-of-distinct-elements-after-operations
func maxDistinctElements(nums []int, k int) int {
	// 计算之后是一个区间，计算区间和数量之间的关系，看最后可以得到多少个数量
	// 将得到的结果尽量放到区间的左边，使用双指针的方式
	sort.Ints(nums)

	ans := 1

	minVal, maxVal := calcRange(nums[0], k)
	ptr := minVal
	for i := 1; i < len(nums); i++ {
		minVal, maxVal = calcRange(nums[i], k)
		// minVal 只有等于 ptr 和 大于 ptr 两种情况
		if minVal > ptr {
			ptr = minVal
			ans++
		} else {
			if maxVal >= ptr+1 {
				ptr = ptr + 1
				ans++
			}
		}
	}

	return ans
}

func calcRange(n, k int) (min, max int) {
	num1 := n + k
	num2 := n - k
	if num1 < num2 {
		return num1, num2
	}

	return num2, num1
}
