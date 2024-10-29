package piscine

func Map(f func(int) bool, a []int) []bool {
	tab := make([]bool, len(a))

	for element1, element2 := range a {
		tab[element1] = f(element2)
	}
	return tab
}
