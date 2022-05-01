package bubble

func sortV1(num []int) {
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
}

func sortV2(num []int) {
	for i := 0; i < len(num); i++ {
		sorted := true
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}

func sortV3(num []int) {
	for end := len(num) - 1; end > 0; end-- {
		// 当数组完全有序时，end-- 结果是0，循环退出
		sortedIndex := 1
		for start := 1; start <= end; start++ {
			if num[start] < num[start-1] {
				num[start-1], num[start] = num[start], num[start-1]
				sortedIndex = start
			}
		}
		end = sortedIndex
	}
}
