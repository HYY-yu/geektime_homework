package 第二周

// 假设自己有个钱包 wall
// 钱包里只有两种类型的钱，5块的，十块的
func lemonadeChange(bills []int) bool {
	wall := make([]int,2)

	for i:= 0;i<len(bills);i++{
		if bills[i] == 5{
			wall[0]++ // 加一张5块
			continue
		}
		if bills[i]==10{
			wall[0]--
			if wall[0] < 0{
				return false
			}
			wall[1]++
		}
		if bills[i] == 20{
			wall[0]--
			if wall[1] == 0{
				wall[0] -=2
			}else if wall[1] > 0{
				wall[1]--
			}
			if wall[0]<0 || wall[1]<0{
				return false
			}
		}
	}
	return true
}
