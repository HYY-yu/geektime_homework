package 第四周

func isPowerrOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(-n) == n
}
