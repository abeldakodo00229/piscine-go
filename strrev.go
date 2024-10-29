package piscine

func StrRev(s string) string {
	a := []rune(s)
	for b, c := 0, len(a)-1; b < c; b, c = b+1, c-1 {
		a[b], a[c] = a[c], a[b]
	}
	return string(a)
}
