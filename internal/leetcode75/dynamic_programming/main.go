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
