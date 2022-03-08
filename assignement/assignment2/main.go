package main

import "fmt"

type area interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	a float64
}

func main() {
	newTriangle := triangle{base: 24, height: 12}
	printArea(newTriangle)

	newSquare := square{a: 10}
	printArea(newSquare)
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (sq square) getArea() float64 {
	return sq.a * sq.a
}

func printArea(a area) {
	fmt.Println("Area: ", a.getArea())
}
