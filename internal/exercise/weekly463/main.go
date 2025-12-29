package weekly463

// 3653. 区间乘法查询后的异或 I
// https://leetcode.cn/problems/xor-after-range-multiplication-queries-i/description/
func xorAfterQueries(nums []int, queries [][]int) int {
	for _, query := range queries {
		idx := query[0]
		for idx <= query[1] {
			nums[idx] = (nums[idx] * query[3]) % (1000000000 + 7)
			idx += query[2]
		}
	}

	ans := 0
	for _, num := range nums {
		ans ^= num
	}

	return ans
}
