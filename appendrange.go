package piscine

func AppendRange(min, max int) []int {
	var sliceArray []int
	for i := min; i < max; i++ {
		sliceArray = append(sliceArray, i)
	}
	return sliceArray
}
