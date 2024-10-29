package piscine

func getInteger3(r rune) int {
	result := 0
	if r == '0' {
		result = 0
	} else if r == '1' {
		result = 1
	} else if r == '2' {
		result = 2
	} else if r == '3' {
		result = 3
	} else if r == '4' {
		result = 4
	} else if r == '5' {
		result = 5
	} else if r == '6' {
		result = 6
	} else if r == '7' {
		result = 7
	} else if r == '8' {
		result = 8
	} else if r == '9' {
		result = 9
	} else {
		result = -1
	}
	return result
}

func Atoi(s string) int {
	tab := []rune(s)
	result := 0
	isValidate := true
	countSign := 0
	for i := 0; i < len(tab); i++ {
		if tab[i] == '-' || tab[i] == '+' {
			countSign++
			if countSign >= 1 && i > 0 {
				isValidate = false
			}
			continue
		}
		if countSign == 0 || countSign == 1 {
			r := getInteger3(tab[i])
			if r != -1 {
				for j := 1; j < len(tab)-i; j++ {
					r = r * 10
				}
			} else {
				isValidate = false
				break
			}
			result += r
		}
	}
	if len(tab) >= 1 {
		if tab[0] == '-' {
			result = result * -1
		}
	}
	if isValidate {
		return result
	} else {
		return 0
	}
}
