package 第三周

import "testing"

// 测试函数
func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}},
			want: 7,
		},
		{
			name: "test2",
			args: args{grid: [][]int{
				{1, 2},
				{5, 6},
				{1, 1},
			}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func minPathSum(grid [][]int) int {
	// DP[i][j] 指二维数组中某元素在此点的 最小路径
	// DP[i][j] = grid[i][j] + MIN(DP[i-1][j],DP[i][j-1])
	// DP[0][0] = grid[0][0]
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	// 吸收大佬解法，不需要dp数组，直接在原grid上进行dp计算
	for i := 0; i < m; i++ {
		// 先特殊处理第一行
		if i == 0 {
			// 特殊处理第一列
			grid[i][0] = grid[i][0]
			for j := 1; j < n; j++ {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}
		} else {
			// 特殊处理第一列
			grid[i][0] = grid[i][0] + grid[i-1][0]
			for j := 1; j < n; j++ {
				grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
