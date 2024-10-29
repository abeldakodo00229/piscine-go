package piscine

func NRune(s string, n int) rune {
	var result rune
	tab := []rune(s)
	if n <= 0 || n > len(tab) {
		return 0
	} else {
		for i, elmt := range s {
			if i == n-1 {
				result = elmt
				break
			}
		}
	}
	return result
}
