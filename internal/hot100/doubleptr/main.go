package doubleptr

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
