func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := mergeSort(s)
	n := mergeSort(t)

	fmt.Println(m, "\n", n)
	return mergeSort(s) == mergeSort(t)
}

func mergeSort(str string) string {
	if len(str) < 2 {
		return str
	}

	middle := len(str) / 2
	left, right := mergeSort(str[:middle]), mergeSort(str[middle:])

	result := make([]string, 0, middle)

	i, j := 0,0
	for i < len(left) && j < len(right) {
		switch {
			case left[i] > right[j]:
				result = append(result, string(right[j]))
				j++
			default:
				result = append(result, string(left[i]))
				i++
		}
	}
	
	s := strings.Join(result, "")
	s += left[i:] 
	s += right[j:]

	return s
}