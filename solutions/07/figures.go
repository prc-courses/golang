package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p *Point) Scale(s float64) {
	p.X = p.X * s
	p.Y = p.Y * s
}

type Rectangle struct {
	Point
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector p1 is: %f\n", p1.Abs())

	p2 := Point{4, 5}
	fmt.Printf("The length of the vector p2 is: %f\n", p2.Abs())

	p1.Scale(5)
	fmt.Printf("The length of the vector p1 is: %f\n", p1.Abs())
	fmt.Printf("Point p1 scaled by 5 has the following coordinates: X %f - Y %f\n", p1.X, p1.Y)

	r1 := new(Rectangle)
	r1.X, r1.Y = 0, 0
	r1.width = 10
	r1.height = 5
	fmt.Printf("The Rectangle r1 area is: %f\n", r1.Area())
	fmt.Printf("The Rectangle r1 perimeter is: %f\n", r1.Perimeter())
}
