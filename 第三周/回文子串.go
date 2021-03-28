package 第三周

// DP 同视频讲解思路一致，把字符串放到矩阵的横、纵
func countSubstrings(s string) int {
	count := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
				count++
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
				count++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				count++
			}
		}
	}
	return count
}
