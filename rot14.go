package piscine

func Rot14(s string) string {
	a := []rune(s)
	tab := []rune{}
	for _, elemt := range a {
		if elemt >= 'a' && elemt <= 'z' {
			tab = append(tab, (elemt-'a'+14)%26+'a')
		} else if elemt >= 'A' && elemt <= 'Z' {
			tab = append(tab, (elemt-'A'+14)%26+'A')
		} else {
			tab = append(tab, elemt)
		}
	}

	return string(tab)
}

/*
func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}*/
