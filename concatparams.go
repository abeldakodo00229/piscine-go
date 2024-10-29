package piscine

func ConcatParams(args []string) string {
	concatArray := ""
	for i := 0; i < len(args); i++ {
		concatArray = concatArray + args[i]
		if i < len(args)-1 {
			concatArray = concatArray + "\n"
		}
	}
	return concatArray
}
