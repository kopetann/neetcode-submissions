func twoSum(nums []int, target int) []int {
	var findSumIndicies func(int) []int
	findSumIndicies = func(idx int) []int {
	if idx >= (len(nums) -1) {
		return []int{}
	}

	decr := target - nums[idx]
	for i, n := range nums {
		if i == idx {
			continue
		}

		if (decr -n) == 0 {
			return []int{idx, i}
		}
	}

	idx++
	return findSumIndicies(idx)
}
	return findSumIndicies(0)
}