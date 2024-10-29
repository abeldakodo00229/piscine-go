package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func Num(nb int) {
	t := '0'
	if nb == 0 {
		z01.PrintRune('0')
		return
	}
	for i := 1; i <= nb%10; i++ {
		t++
	}
	for i := -1; i >= nb%10; i-- {
		t++
	}
	if nb/10 != 0 {
		Num(nb / 10)
	}
	z01.PrintRune(t)
}

func main() {
	points := &point{}

	setPoint(points)

	// z01.PrintRune("x = %d, y = %d\n", points.x, points.y)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Num(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Num(points.y)
	z01.PrintRune('\n')
}
