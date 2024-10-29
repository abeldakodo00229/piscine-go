package piscine

func ToUpper(s string) string {
	lowerCase := []rune{}
	upperCase := []rune{}
	upperResult := []rune{}

	for i := 65; i <= 90; i++ {
		upperCase = append(upperCase, int32(i))
	}
	for i := 97; i <= 122; i++ {
		lowerCase = append(lowerCase, int32(i))
	}
	for _, elemtn := range s {
		found := 0
		i := 0
		for i = 0; i < len(lowerCase); i++ {
			if elemtn == lowerCase[i] {
				found++
				break
			}
		}
		if found == 1 {
			upperResult = append(upperResult, int32(upperCase[i]))
		} else {
			upperResult = append(upperResult, elemtn)
		}
	}
	return string(upperResult)
}
