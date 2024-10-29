package piscine

func getNumbers(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result = result * 10
	}
	return result
}

func TrimAtoi(s string) int {
	aphaCase := []rune{}
	resultArray := []rune{}
	for i := 48; i <= 57; i++ {
		aphaCase = append(aphaCase, int32(i))
	}
	aphaCase = append(aphaCase, 45)
	for _, elemtn := range s {
		found := 0
		for i := 0; i < len(aphaCase); i++ {
			if elemtn == aphaCase[i] {
				found++
				break
			}
		}
		if found == 1 {
			resultArray = append(resultArray, elemtn)
		}
	}
	if len(resultArray) == 0 {
		return 0
	} else {
		intResult := 0
		numericTab := []rune{}
		for i := 0; i < len(resultArray); i++ {
			if resultArray[i] != 45 {
				numericTab = append(numericTab, resultArray[i])
			}
		}
		for i := 0; i < len(numericTab); i++ {
			if numericTab[i] != 45 {
				intResult = intResult + int(numericTab[i]-48)*getNumbers(len(numericTab)-1-i)
			}
		}
		if resultArray[0] == 45 {
			intResult = -1 * intResult
		}
		return intResult
	}
}
