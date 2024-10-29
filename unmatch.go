package piscine

/*import (
	"fmt"
)*/

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				a[i] = -1
				a[j] = -1

			}
		}
	}
	for k := 0; k < len(a); k++ {
		if a[k] != -1 {
			return a[k]
		}
	}
	return -1
}

/*
func main() {
	arr := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(arr)
	fmt.Println(unmatch)
}
*/
