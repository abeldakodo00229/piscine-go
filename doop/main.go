package main

import (
	"os"
)

func bacatoi(s int) string {
	tab := []rune{}
	for s > 0 {
		tab = append(tab, rune(s%10+48))
		s = s / 10
	}
	for i := 0; i < len(tab); i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] < tab[j] {
				c := tab[j]
				tab[j] = tab[i]
				tab[i] = c
			} else if tab[i] > tab[j] {
				c := tab[j]
				tab[j] = tab[i]
				tab[i] = c
			}
		}
	}
	return string(tab)
}

func basicamoi(s string) int {
	c := 0
	p := 1
	if s[0] == '-' {
		p = -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			c = c*10 + int(s[i]-'0')
		}
	}
	return c * p
}

func afficherstr(s string) {
	b := []byte{}
	for _, a := range s {
		b = append(b, byte(a))
	}
	b = append(b, '\n')
	os.Stdout.Write(b)
}

func calcul(a int, b int, o string) {
	s := "No division by 0"
	s1 := "No modulo by 0"
	res := 0
	p := 1
	if a >= -2147483648 && a <= 2147483647 && b >= -2147483648 && b <= 2147483647 {
		if o == "*" {
			res = a * b
			p = 0
		}
		if o == "+" {
			res = a + b
			p = 0
		}
		if o == "-" {
			res = a - b
			p = 0
		}
		if o == "/" {
			if a >= 0 && b != 0 {
				res = a / b
				p = 0
			} else if a != 0 && b == 0 {
				afficherstr(s)
			}
		}
		if o == "%" {
			if a >= 0 && b != 0 {
				res = a % b
				p = 0
			} else if a != 0 && b == 0 {
				afficherstr(s1)
			}
		}
		if p == 0 {
			if res < 0 {
				os.Stdout.Write([]byte("-"))
				res = -1 * res
				afficherstr(bacatoi(res))
			} else if res == 0 {
				afficherstr("0")
			} else {
				afficherstr(bacatoi(res))
			}
		}
	}
}

func main() {
	arg := os.Args
	p := 0
	t := 0
	if len(arg) == 4 {
		arg := arg[1:]
		for i := 0; i < len(arg[0]); i++ {
			if arg[0][i] >= '0' && arg[0][i] <= '9' {
				p = 1
			}
		}
		for i := 0; i < len(arg[2]); i++ {
			if arg[2][i] >= '0' && arg[2][i] <= '9' {
				t = 1
			}
		}
		if t+p == 2 {
			calcul(basicamoi(arg[0]), basicamoi(arg[2]), arg[1])
		}
	}
}
