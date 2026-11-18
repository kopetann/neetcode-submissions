func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int, len(s))
	for _, ch := range s{
		m[ch] = m[ch]+1
	}

	for _, ch := range t{
		if val, exists := m[ch]; !exists || val < 1 {
			return false
		}

		m[ch]=m[ch]-1
	}

	return true
}
