package algos

func (s *AlgosService) FindSubstring(str string) string {
	maxStr := ""
	left := 0
	charToIndex := make(map[byte]int)

	for right := 0; right < len(str); right++ {
		if ind, ok := charToIndex[str[right]]; ok && ind >= left {
			left = ind + 1
		}

		if len(maxStr) <= right-left {
			maxStr = str[left : right+1]
		}

		charToIndex[str[right]] = right
	}
	return maxStr
}
