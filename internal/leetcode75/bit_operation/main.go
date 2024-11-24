package main

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number
func singleNumber(nums []int) int {
	res := 0
	// 异或运算有个重要的性质，两个相同数字异或为 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
