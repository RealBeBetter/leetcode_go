package main

// 70. 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/description
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// 单纯的递归会超时
	return climbStairs(n-1) + climbStairs(n-2)
}

// 70. 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/description
func climbStairsWithCache(n int) int {
	if n <= 2 {
		return n
	}

	pre1, pre2 := 1, 2
	var res int

	for i := 3; i <= n; i++ {
		res = pre1 + pre2
		pre1 = pre2
		pre2 = res
	}

	return res
}

// 118. 杨辉三角
// https://leetcode.cn/problems/pascals-triangle/description
func generate(numRows int) [][]int {
	res := [][]int{{1}}
	if numRows == 1 {
		return res
	}

	res = append(res, []int{1, 1})

	for i := 3; i <= numRows; i++ {
		row := make([]int, i)
		// 处理首位
		row[0], row[len(row)-1] = 1, 1
		// 获取上一行
		preRow := res[len(res)-1]
		for j := 1; j < len(row)-1; j++ {
			row[j] = preRow[j] + preRow[j-1]
		}

		res = append(res, row)
	}

	return res
}
