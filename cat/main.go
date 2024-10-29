package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	i := 1
	if len(os.Args) > 1 {
		for i < len(os.Args) {
			file, err := ioutil.ReadFile(os.Args[i])

			if err != nil {
				no_directory_error := "ERROR: " + err.Error() + "\n"
				PrintStr(no_directory_error)
				os.Exit(1)
			} else {
				PrintStr(string(file))
			}
			i++
		}
	} else {
		file, _ := ioutil.ReadAll(os.Stdin)
		PrintStr(string(file))
	}
}
