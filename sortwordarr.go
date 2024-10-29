package piscine

func SortWordArr(array []string) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if Compare(array[i], array[j]) > 0 {
				ech := array[i]
				array[i] = array[j]
				array[j] = ech
			}
		}
	}
}
