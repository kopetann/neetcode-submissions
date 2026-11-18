func hasDuplicate(nums []int) bool {
    buf := make([]int, len(nums))

	return mergeSort(nums, buf)
}

func mergeSort(nums, buf []int) bool {
	if len(nums) < 2 {
		return false
	}

	mid := len(nums) / 2

	if mergeSort(nums[:mid], buf[:mid]) || mergeSort(nums[mid:], buf[mid:]) {
		return true
	}

	copy(buf, nums)
	left, right := buf[:mid], buf[mid:len(nums)]

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		switch {
			case left[i] == right[j]:
				fmt.Printf("left %d \t right %d", left[i], right[j])
				return true
			case left[i] < right[j]:
				nums[k] = left[i]
				i++
			case right[j] < left[i]:
				nums[k] = right[j]
				j++
		}
		k++
	}

	k += copy(nums[k:], left[i:])
	copy(nums[k:], right[j:])

	return false
}