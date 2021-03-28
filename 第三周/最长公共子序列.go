package 第三周

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0

	}
	pre, curr := 1, 1
	for i := 1; i < len(s); i++ {
		p := curr
		if s[i] == '0' {
			// s[i] == 0 的情况
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
				continue
			}
			return 0
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			curr += pre
		}
		pre = p
	}
	return curr
}
