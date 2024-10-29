package piscine

// import "fmt"

func Abort(a, b, c, d, e int) int {
	t := []int{a, b, c, d, e}
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t); j++ {
			if t[i] > t[j] {
				t[i], t[j] = t[j], t[i]
			}
		}
	}
	return t[len(t)/2]
}

/*
func main() {
	middle := Abort(-852200979910544858, 133029217788333556, 4727557918776213580, -6544010413339254750, 4944718688999527423)
	fmt.Println(middle)
}*/
