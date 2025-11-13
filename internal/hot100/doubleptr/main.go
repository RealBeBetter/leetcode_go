package doubleptr

import (
	"sort"
)

// 283. 移动零
// https://leetcode.cn/problems/move-zeroes/description/
func moveZeroes(nums []int) {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		nums[idx] = nums[i]
		idx++
	}

	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 283. 移动零
// https://leetcode.cn/problems/move-zeroes/description/
func moveZerosWithPtr(nums []int) {
	ptr := 0
	for i, x := range nums {
		if x != 0 {
			// 维护一个全是零的区间
			nums[i], nums[ptr] = nums[ptr], x
			ptr++
		}
	}
}

// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/description
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := min(height[l], height[r]) * (r - l)
	// l++ 与 r-- 都是为了获得更多的水，如果短的木板不移动，永远无法获取比上次更多的水
	for l < r {
		temp := min(height[l], height[r]) * (r - l)
		res = max(res, temp)

		if height[l] > height[r] {
			r--
			continue
		}

		l++
	}

	return res
}

// 15. 三数之和
// https://leetcode.cn/problems/3sum/description
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// 根据双指针的移动方向，排除重复的答案
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				k--
				j++
			}
		}
	}

	return res
}

// 560. 和为 K 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k/description
func subarraySum(nums []int, k int) int {
	prefixSum := 0
	prefixSumMap := map[int]int{0: 1}
	res := 0

	for _, num := range nums {
		prefixSum += num

		if cnt, ok := prefixSumMap[prefixSum-k]; ok {
			res += cnt
		}

		prefixSumMap[prefixSum]++
	}

	return res
}

// 560. 和为 K 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k/description
func subarraySumWithDoublePtr(nums []int, k int) int {
	// 时间复杂度比较高
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}

	return res
}

// 560. 和为 K 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k/description
func subarraySumWithDoublePtrReverse(nums []int, k int) int {
	// 时间复杂度比较高
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}

	return res
}

// 57. 插入区间
// https://leetcode.cn/problems/insert-interval/description/
func insert(intervals [][]int, newInterval []int) [][]int {
	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	minInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := make([][]int, 0, len(intervals))

	newStart := newInterval[0]
	newEnd := newInterval[1]

	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		// 新区间在左边，已经完成合并，从当前节点返回
		if newEnd < start {
			res = append(res, []int{newStart, newEnd})
			res = append(res, intervals[i:]...)
			return res
		} else if newStart > end {
			// 新区间在右边，添加当前区间
			res = append(res, []int{start, end})
		} else {
			// 新区间有重叠，继续计算下一个区间
			newStart = minInt(start, newStart)
			newEnd = maxInt(end, newEnd)
		}
	}

	res = append(res, []int{newStart, newEnd})
	return res
}
