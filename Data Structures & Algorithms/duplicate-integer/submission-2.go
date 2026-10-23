func hasDuplicate(nums []int) bool {
	m := make(map[int]int, len(nums))

	for idx, num := range nums {
		if sudIdx, exists := m[num]; exists && idx != sudIdx {
			return true
		}

		m[num] = idx 
	}

	return false
}
