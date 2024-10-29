package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 1 {
		return nb
	} else if power == 0 {
		return 1
	} else {
		a := 1
		for i := 0; i < power; i++ {
			a *= nb
		}
		return a
	}
}
