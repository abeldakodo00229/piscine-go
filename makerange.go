package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	sliceArray := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		sliceArray[i] = min + i
	}
	return sliceArray
}
