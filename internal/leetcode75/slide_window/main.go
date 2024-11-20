package slide_window

// 643. 子数组最大平均数 I
// https://leetcode.cn/problems/maximum-average-subarray-i/
func findMaxAverage(nums []int, k int) float64 {
	var res float64
	var lastTemp float64

	for i := 0; i < k; i++ {
		lastTemp += float64(nums[i])
	}

	divisor := float64(k)
	res = lastTemp / divisor

	for i := k; i < len(nums); i++ {
		lastTemp = lastTemp - float64(nums[i-k]) + float64(nums[i])

		avg := lastTemp / divisor
		if avg > res {
			res = avg
		}
	}

	return res
}
