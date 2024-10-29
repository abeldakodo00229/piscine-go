package piscine

// import "github.com/01-edu/z01"

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 1234 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
