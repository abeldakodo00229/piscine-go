package piscine

/*import (
	"fmt"
)
*/
func ActiveBits(n int) int {
	a := []int{}
	if n < 0 {
		n = -n
	}
	for n > 0 {
		a = append(a, n%2)
		n = n / 2
	}
	count := 0
	for _, elmt := range a {
		if elmt == 1 {
			count++
		}
	}
	return count
}

/*
func main() {
	nbits := ActiveBits(16)
	fmt.Println(nbits)
}*/
