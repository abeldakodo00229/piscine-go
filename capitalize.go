package piscine

func Capitalize(s string) string {
	lowerCase := []rune{}
	upperCase := []rune{}
	aphaCase := []rune{}
	arrayList := []rune(s)
	// ALL APHA NUMERIC
	for i := 32; i <= 126; i++ {
		aphaCase = append(aphaCase, int32(i))
	}
	// UPPER
	for i := 65; i <= 90; i++ {
		upperCase = append(upperCase, int32(i))
	}
	// LOWER
	for i := 97; i <= 122; i++ {
		lowerCase = append(lowerCase, int32(i))
	}
	for index, element := range s {
		if IsAlpha(string(element)) {
			if index == 0 || !IsAlpha(string(arrayList[index-1])) {
				if IsLower(string(element)) {
					result := []rune(ToUpper(string(arrayList[index])))
					arrayList[index] = result[0]
				}
			} else {
				if IsUpper(string(element)) {
					result := []rune(ToLower(string(arrayList[index])))
					arrayList[index] = result[0]
				}
			}
		}
	}
	return string(arrayList)
}
