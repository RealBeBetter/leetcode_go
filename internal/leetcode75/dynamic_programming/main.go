package main

// 790. 多米诺和托米诺平铺
// https://leetcode.cn/problems/domino-and-tromino-tiling
func numTilings(n int) int {
	// f(i) = f(i-1) + f(i-2) + 2 (f(i-3) + ... + 1)
	// 简化得到：f(n) = 2*f(n-1) + f(n-3)
	dp := make([]int, n+1)
	if n == 1 {
		return 1
	}

	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		nextVal := 2*dp[i-1] + dp[i-3]
		dp[i] = nextVal % (1e9 + 7)
	}

	return dp[n]
}

// 72. 编辑距离
// https://leetcode.cn/problems/edit-distance/description
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for i := 0; i <= len(word2); i++ {
		dp[0][i] = i
	}

	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j], dp[i][j]) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
