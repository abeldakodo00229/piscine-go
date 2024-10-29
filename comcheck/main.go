package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) >= 2 {
		tab := arg[1:]
		for i := 0; i < len(tab); i++ {
			if tab[i] == "01" || tab[i] == "galaxy" || tab[i] == "galaxy 01" {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
