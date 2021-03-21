package 第二周

import "sort"

func findContentChildren(greed, size []int) (ans int) {
	sort.Ints(greed)
	sort.Ints(size)
	lenGreed, lenSize := len(greed), len(size)
	for i, j := 0, 0; i < lenGreed && j < lenSize; i++ {
		// 找个能满足胃口最大的孩子的饼干
		for ; j < lenSize && greed[i] > size[j]; j++ {
		}
		// 找到这个饼干就分配给他
		if j < lenSize {
			ans++
			j++
		}
	}
	return
}
