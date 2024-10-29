package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	array := []rune{}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for n > 0 {
		array = append(array, int32(n%10))
		n = n / 10
	}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				tmp := array[j]
				array[j] = array[i]
				array[i] = tmp
			}
		}
	}
	for i := 0; i < len(array); i++ {
		z01.PrintRune(array[i] + 48)
	}
}
