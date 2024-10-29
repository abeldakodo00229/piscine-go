package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	compteurBase := 0
	for _, a1 := range base {
		differentCaracter := 0
		if a1 == '+' || a1 == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for _, a2 := range base {
			if a1 == a2 {
				differentCaracter++
			}
			if differentCaracter == 2 {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
		compteurBase++
	}
	if compteurBase < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	mod := ""
	va := []rune(base)
	if nbr < 0 {
		z01.PrintRune('-')
	}
	for ; nbr != 0; nbr /= compteurBase {
		index := nbr % compteurBase
		if index < 0 {
			index = -index
		}
		mod = string(va[index]) + mod
	}
	for _, r := range mod {
		z01.PrintRune(r)
	}
}
