package greedy

// 55. 跳跃游戏
// https://leetcode.cn/problems/jump-game/description/
func canJump(nums []int) bool {
	maxRange := nums[0]
	for i, num := range nums {
		if i > maxRange {
			return false
		}
		maxRange = max(maxRange, i+num)
	}

	return maxRange >= len(nums)-1
}

// 45. 跳跃游戏 II
// https://leetcode.cn/problems/jump-game-ii/description/
func jump(nums []int) int {
	steps := 0

	boundary := 0
	farthest := 0

	// 最终到达最后一个元素，实际上可以不访问最后一个元素
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		// 到达当前这次跳跃的边界值，更新边界值为下一次跳跃的最远距离，发生一次真实跳跃
		if i == boundary {
			boundary = farthest
			steps++
		}
	}

	return steps
}
