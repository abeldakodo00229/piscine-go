package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		for _, element := range a[i] {
			z01.PrintRune(element)
		}
		z01.PrintRune('\n')
	}
}
