func twoSum(nums []int, target int) []int {
    i, j := 0,0
	for i < len(nums) {
		j++
		if i == j {
			continue
		}

		switch {
			case nums[i] + nums[j] == target:
				return []int{i, j}
			case j == (len(nums) - 1):
				i++
				j=0
		}
	}

	return []int{}
}
