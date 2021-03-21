package 第二周

// 1. hash表记录
func majorityElement(nums []int) int {
	hash := make(map[int]int)
	n := len(nums)

	for i := 0; i < n; i++ {
		hash[nums[i]]++
		if hash[nums[i]] > n/2 {
			return nums[i]
		}
	}
	return 0
}


// 2. 分治法
func majorityElementF(nums []int) int {
	return mF(nums, 0, len(nums)-1)
}

func mF(nums []int, s, e int) int {
	// 终止条件
	if s == e {
		return nums[s]
	}

	mid := (e-s)/2 + s
	lnum := mF(nums, s, mid)
	rnum := mF(nums, mid+1, e)

	if lnum == rnum {
		return lnum
	}
	// 合并策略
	lcount := countNums(nums, lnum, s, e)
	rcount := countNums(nums, rnum, s, e)

	if lcount > rcount {
		return lnum
	}
	return rnum
}

func countNums(nums []int, num int, s, e int) int {
	count := 0
	for i := s; i <= e; i++ {
		if nums[i] == num {
			count++
		}
	}
	return count
}

// 3. 摩尔投票法
func majorityElementMore(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cand := nums[0]
	count := 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			cand = nums[i]
		}
		if cand == nums[i] {
			count++
		} else {
			count--
		}
	}

	// 若无法保证必然存在 多余n/2 长度的元素，在选出cand后应该进行校验
	return cand
}