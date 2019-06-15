package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{height: 3, base: 4}
	s := square{sideLength: 4}

	fmt.Println("Triangle area is:", printArea(t))
	fmt.Println("Square area is:", printArea(s))
}

func printArea(s shape) float64 {
	return s.getArea()
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
