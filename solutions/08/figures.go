package main

import "fmt"

type Square struct {
	side float32
}

type Triangle struct {
	base, height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	Perimeter() float32
}

type Figures interface{}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return sq.side * 4
}

func (tr *Triangle) Area() float32 {
	return tr.base * tr.height / 2
}

func main() {
	sq1 := &Square{10}
	sq2 := &Square{15}
	tr := &Triangle{15, 10}

	figures := []Figures{sq1, sq2, tr}

	for _, f := range figures {
		if v, ok := f.(AreaInterface); ok {
			fmt.Println("Area of this figure is:", v.Area())
		}
		if v, ok := f.(PeriInterface); ok {
			fmt.Println("Perimeter of this figure is:", v.Perimeter())
		}
	}
}
