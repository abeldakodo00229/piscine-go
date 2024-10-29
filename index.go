package piscine

func Index(s string, toFind string) int {
	sArray := []rune(s)
	sArrayFind := []rune(toFind)
	for i := 0; i <= len(sArray)-len(sArrayFind); i++ {
		j := 0
		for j = 0; j < len(sArrayFind); j++ {
			if sArray[i+j] != sArrayFind[j] {
				break
			}
		}
		if j == len(sArrayFind) {
			return i
		}
	}
	return -1
}
