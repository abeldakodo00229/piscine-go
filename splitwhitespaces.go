package piscine

func SplitWhiteSpaces(s string) []string {
	globalArray := []rune(s)
	array := []string{}
	previous := 0
	for index, element := range s {
		if index == 0 && element == ' ' {
			previous = 1
		} else {
			if element == ' ' || index == len(globalArray)-1 {
				localArray := []rune{}
				max := index
				if index == len(globalArray)-1 {
					max = index + 1
				}
				for i := previous; i < max; i++ {
					if globalArray[i] != ' ' {
						localArray = append(localArray, globalArray[i])
					}
				}
				array = append(array, string(localArray))
				previous = index + 1
			}
		}
	}
	result := []string{}
	for i := 0; i < len(array); i++ {
		if array[i] != "" {
			result = append(result, array[i])
		}
	}
	return result
}
