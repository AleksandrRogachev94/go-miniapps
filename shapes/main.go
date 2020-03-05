package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (sq square) area() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) area() float64 {
	return tr.height * tr.base * 0.5
}

func printArea(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	sq := square{sideLength: 5.5}
	tr := triangle{height: 3, base: 5}

	printArea(sq)
	printArea(tr)
}
