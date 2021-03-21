package 第二周

// 贪心
// 只要下一天是涨的，我就买入，且在第二天卖出
func maxProfit2(prices []int) int {
	// 7,1,5,3,6,4
	sum := 0
	pre := false // 指示前一天是否卖出
	for i := 0; i < len(prices); i++ {
		next := 0
		if i >= len(prices)-1 {
			next = 0
		} else {
			next = prices[i+1]
		}
		// 卖出
		if pre {
			sum += prices[i]
			pre = false
		}
		if next > prices[i] {
			// 买入
			sum -= prices[i]
			pre = true
		}
	}
	return sum
}

