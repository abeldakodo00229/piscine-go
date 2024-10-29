package piscine

func AlphaCount(s string) int {
	aphaCase := []int{}

	for i := 65; i <= 90; i++ {
		aphaCase = append(aphaCase, i)
	}

	for i := 97; i <= 122; i++ {
		aphaCase = append(aphaCase, i)
	}

	count := 0

	for _, element := range s {
		for i := 0; i < len(aphaCase); i++ {
			if rune(aphaCase[i]) == element {
				count++
			}
		}
	}

	return count
}
