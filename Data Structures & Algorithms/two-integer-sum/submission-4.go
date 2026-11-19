func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
	
	for i, n := range nums {
		v := target - n
		if idx, exists := m[v]; exists {
			return []int{idx, i}
		}

		m[n]=i
	}

	return []int{}
}
