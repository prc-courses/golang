package main

type Square struct {
	side float32
}

type Triangle struct {
	// Your code here
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	// Your code here
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	// Your code here
	return 0
}

func (tr *Triangle) Area() float32 {
	// Your code here
	return 0
}

func main() {
	// Your code here
}
