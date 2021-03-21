package 第二周

// 50. Pwo(x,n)

// 分治法
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickPow(x, n)
	}
	return 1 / quickPow(x, -n)
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	subResult := quickPow(x, n)
	if n%2 == 0 {
		return subResult * subResult
	}
	return subResult * subResult * x
}
