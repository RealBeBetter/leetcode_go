package main

// 221. 最大正方形
// https://leetcode.cn/problems/maximal-square/description
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	maxVal := 0

	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			val := int(matrix[i][j] - '0')
			dp[i][j] = val
			if val == 1 {
				maxVal = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] != 1 {
				continue
			}

			val := min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			dp[i][j] = val
			maxVal = max(maxVal, val)
		}
	}

	return maxVal * maxVal
}
