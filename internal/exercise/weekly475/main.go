package weekly475

import "math"

// 3740. 三个相等元素之间的最小距离 I
// https://leetcode.cn/problems/minimum-distance-between-three-equal-elements-i/description/
func minimumDistance(nums []int) int {
	minDistance := math.MaxInt
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[j] == nums[k] {
					minDistance = min(minDistance, 2*(k-i))
				}
			}
		}
	}

	if minDistance == math.MaxInt {
		return -1
	}
	return minDistance
}

// 3741. 三个相等元素之间的最小距离 II
// https://leetcode.cn/problems/minimum-distance-between-three-equal-elements-ii/
func minimumDistanceII(nums []int) int {
	m := make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}

	minDistance := math.MaxInt
	for _, indexes := range m {
		// 因为添加时的 idx 是递增的，所以直接按序计算即可
		for i := 2; i < len(indexes); i++ {
			minDistance = min(minDistance, 2*(indexes[i]-indexes[i-2]))
		}
	}

	if minDistance == math.MaxInt {
		return -1
	}
	return minDistance
}
