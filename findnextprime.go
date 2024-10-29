package piscine

func IsLocalPrime(nb int) bool {
	if nb == 2 || nb == 3 {
		return true
	}
	if nb <= 1 || nb%2 == 0 || nb%3 == 0 {
		return false
	}
	for i := 5; i*i <= nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	} else {
		nextPrime := nb
		for nb <= nextPrime {
			if IsLocalPrime(nextPrime) {
				break
			} else {
				nextPrime++
			}
		}
		return nextPrime
	}
}
