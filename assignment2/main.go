package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	t1 := triangle{
		base:   3,
		height: 3,
	}

	s1 := square{
		sideLength: 5,
	}

	printArea(t1)
	printArea(s1)
}

func printArea(s shape) {
	fmt.Println("The area is:", s.getArea())
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
