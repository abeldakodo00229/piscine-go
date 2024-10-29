package piscine

func Compact(ptr *[]string) int {
	a := 0
	var tabstr []string
	for _, elmt := range *ptr {
		if elmt != "" {
			tabstr = append(tabstr, elmt)

			a++
		}
	}
	*ptr = tabstr
	return a
}

/*
const N = 6

func main() {
	arr := make([]string, N)
	arr[0] = "a"
	arr[2] = "b"
	arr[4] = "c"

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&arr))

	for _, v := range arr {
		fmt.Println(v)
	}
}*/
