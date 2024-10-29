package piscine

func IsAlpha(s string) bool {
	aphaCase := []rune{}
	arrayList := []rune(s)
	for i := 65; i <= 90; i++ {
		aphaCase = append(aphaCase, int32(i))
	}
	for i := 97; i <= 122; i++ {
		aphaCase = append(aphaCase, int32(i))
	}
	for i := 48; i <= 57; i++ {
		aphaCase = append(aphaCase, int32(i))
	}
	found := 0
	for _, elemtn := range s {
		for a := 0; a < len(aphaCase); a++ {
			if elemtn == aphaCase[a] {
				found++
				break
			}
		}
	}
	if found == len(arrayList) {
		return true
	} else {
		return false
	}
}
