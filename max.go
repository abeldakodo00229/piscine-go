package piscine

// import "fmt"

func Max(a []int) int {
	t := a
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t); j++ {
			if t[i] > t[j] {
				t[i], t[j] = t[j], t[i]
			}
		}
	}
	return t[0]
}

/*
func main() {
	arrInt := []int{23, 123, 1, 1100, 55, 93}
	max := Max(arrInt)
	fmt.Println(max)
}*/
