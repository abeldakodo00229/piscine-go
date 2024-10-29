package piscine

// import "fmt"

func CollatzCountdown(start int) int {
	count := 0
	if start <= 0 {
		return -1
	}
	if start > 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = 3*start + 1
		}
		count = CollatzCountdown(start)
		count++
	}
	return count
}

/*
func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}
*/
