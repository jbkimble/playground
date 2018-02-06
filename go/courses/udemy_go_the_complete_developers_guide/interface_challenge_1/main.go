package main

import (
	"fmt"
	// "strconv"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength float64
}

func main() {
	tr := triangle{height: 10.0, base: 11.0}
	sq := square{sideLength: 15.0}

	printArea(tr)
	printArea(sq)
}

func (tr triangle) getArea() float64 {
	totalArea := 0.5 * tr.height * tr.base
	return totalArea
}

func (sq square) getArea() float64 {
	totalArea := sq.sideLength * sq.sideLength
	return totalArea
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}