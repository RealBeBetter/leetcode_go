package main

import (
	"math"
)

// 704. 二分查找
// https://leetcode.cn/problems/binary-search
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

// 27. 移除元素
// https://leetcode.cn/problems/remove-element
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}

		// 替换
		nums[idx] = nums[i]
		idx++
	}

	return idx
}

// 977. 有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array
func sortedSquares(nums []int) []int {
	// 返回平方后的数组
	if len(nums) == 0 {
		return nums
	}

	// 从新数组的末尾开始写入
	idx := len(nums) - 1
	res := make([]int, len(nums))

	left, right := 0, len(nums)-1
	// 双指针，从两端分别取值，到最后在中间的位置相遇
	for left <= right {
		if abs(nums[left]) <= abs(nums[right]) {
			res[idx] = nums[right] * nums[right]
			right--
		} else if abs(nums[left]) > abs(nums[right]) {
			res[idx] = nums[left] * nums[left]
			left++
		}

		idx--
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

// 209. 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	sum := 0

	start := 0
	// 双指针，固定结束的位置
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 循环，找到最短的子数组
		for sum >= target {
			tempLen := i - start + 1
			if tempLen < res {
				res = tempLen
			}
			sum -= nums[start]
			start++
		}
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}

// 59. 螺旋矩阵 II
// https://leetcode.cn/problems/spiral-matrix-ii
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	times := n * n
	direction := right

	x, y := 0, 0
	for i := 1; i <= times; i++ {
		matrix[x][y] = i
		x, y, direction = calcNextIdx(x, y, matrix, direction)
	}

	return matrix
}

const (
	right = iota
	down
	left
	up
)

// 计算下一个填充的坐标
func calcNextIdx(x, y int, matrix [][]int, direction int) (nextX, nextY, nextDirection int) {
	// 方向优先级，依次是：向右、向下、向左、向上
	n := len(matrix)

	// 根据指定方向上是否符合边界条件、且是否未被填充；否则就换方向
	if direction == right {
		if y < n-1 && matrix[x][y+1] == 0 {
			return x, y + 1, right
		} else {
			return x + 1, y, down
		}
	}

	if direction == down {
		if x < n-1 && matrix[x+1][y] == 0 {
			return x + 1, y, down
		} else {
			return x, y - 1, left
		}
	}

	if direction == left {
		if y > 0 && matrix[x][y-1] == 0 {
			return x, y - 1, left
		} else {
			return x - 1, y, up
		}
	}

	if direction == up {
		if x > 0 && matrix[x-1][y] == 0 {
			return x - 1, y, up
		} else {
			return x, y + 1, right
		}
	}

	return x, y, direction
}

// 59. 螺旋矩阵 II
// https://leetcode.cn/problems/spiral-matrix-ii
func generateMatrixWithLoop(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	times := n / 2
	count := 1
	offset := 1
	startX, startY := 0, 0
	for i := 0; i < times; i++ {
		nextX, nextY := startX, startY
		// 从左到右
		for ; nextY < n-offset; nextY++ {
			matrix[nextX][nextY] = count
			count++
		}

		// 从上到下
		for ; nextX < n-offset; nextX++ {
			matrix[nextX][nextY] = count
			count++
		}

		// 从右到左
		for ; nextY > startY; nextY-- {
			matrix[nextX][nextY] = count
			count++
		}

		// 从下到上
		for ; nextX > startX; nextX-- {
			matrix[nextX][nextY] = count
			count++
		}

		startX++
		startY++

		offset++
	}

	if n%2 == 1 {
		mid := n / 2
		matrix[mid][mid] = count
	}

	return matrix
}
