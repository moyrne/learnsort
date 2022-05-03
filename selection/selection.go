package selection

func sortV1(nums []int) {
	for end := len(nums) - 1; end > 0; end-- {
		// maxIdx 默认为 0, 所以会访问到 0
		var maxIdx int
		for start := 1; start <= end; start++ {
			if nums[maxIdx] < nums[start] {
				maxIdx = start
			}
		}
		nums[end], nums[maxIdx] = nums[maxIdx], nums[end]
	}
}
