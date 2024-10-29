package piscine

func Swap(a *int, b *int) {
	var c int = *a
	var d int = *b
	*a = d
	*b = c
}
