package piscine

// import "fmt"

func Join(strs []string, sep string) string {
	strprint := ""
	for i := 0; i < len(strs); i++ {
		strprint = strprint + strs[i]
		if i < len(strs)-1 {
			strprint = strprint + sep
		}
	}
	return strprint
}

/*
func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
*/
