package piscine

func IsPrime(nb int) bool {
	count := 0
	if nb <= 1 {
		return false
	} else {
		for i := 1; i <= nb; i++ {
			if nb%i == 0 {
				count++
				if count == 3 {
					break
				}
			}
		}
		if count == 2 {
			return true
		} else {
			return false
		}
	}
}
