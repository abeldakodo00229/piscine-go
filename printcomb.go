package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var (
		i rune
		j rune
		k rune
	)
 
	for i = '0'; i <= '9'; i++ {
		for j = '1'; j <= '9'; j++ {
			for k = '2'; k <= '9'; k++ {
             
				if  i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
            if i =='7' && j =='8' && k=='9'{
					
              }else{
					z01.PrintRune(',')
              z01.PrintRune(' ')
              }

				}
			}
		}
	}
	z01.PrintRune('\n')
}


func main() {
	PrintComb()
}