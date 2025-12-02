package weekly476

import (
	"sort"
	"strings"
)

// 3745. 三元素表达式的最大值
// https://leetcode.cn/problems/maximize-expression-of-three-elements/
func maximizeExpressionOfThree(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1] + nums[len(nums)-2] - nums[0]
}

// 3746. 等量移除后的字符串最小长度
// https://leetcode.cn/problems/minimum-string-length-after-balanced-removals/
func minLengthAfterRemovals(s string) int {
	// 实际上的最小长度，就是 a、b 数量的差值
	cnt := strings.Count(s, "a")
	anotherCnt := len(s) - cnt

	minLen := cnt - anotherCnt
	if minLen < 0 {
		return -minLen
	}
	return minLen
}
