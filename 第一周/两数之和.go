package week1

// 暴力遍历+hash表存结果
func twoSumHash(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := hash[target-nums[i]]; ok {
			return []int{v, i}
		}
		hash[nums[i]] = i
	}
	return nil
}
