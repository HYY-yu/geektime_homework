package 第三周

// 本题没能想到好的思路，仿照官方写了一个
func leastInterval(tasks []byte, n int) int {
	cnt := map[byte]int{}
	// 初始化
	for _, t := range tasks {
		cnt[t]++
	}

	maxExec, maxExecCnt := 0, 0
	for _, c := range cnt {
		if c > maxExec {
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			maxExecCnt++
		}
	}

	if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
		return time
	}
	return len(tasks)
}
