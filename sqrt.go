package piscine

// import "fmt"

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb == 2 {
		return 0
	}
	if nb > 0 {
		result := 1
		sqrt := 0
		for a := 1; result <= nb; a++ {
			result = a * a
			sqrt++
			if result == nb {
				return sqrt
			}
		}
		return 0
	} else {
		return 0
	}
}

/*func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(0))
}*/
