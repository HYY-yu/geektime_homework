package 第三周

// Mark 一种解法
// 动规难的是子问题拆解，和识别这个问题是否可以动规求解
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	// ans 收集最大边
	ans, dp := 0, make([][]int, len(matrix))
	for i, m, n := 0, len(matrix), len(matrix[0]); i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans * ans
}

func min3(x, y, z int) int {
	if x > y {
		x = y
	}
	if x > z {
		return z
	}
	return x
}
