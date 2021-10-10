package main

import (
	"fmt"
	"testing"
)

var solvertests = []struct {
	in  polar
	out string
}{
	{polar{1, 90}, "0.00 1.00"},
	{polar{35, 45}, "24.75 24.75"},
	{polar{30, 30}, "25.98 15.00"},
	{polar{60, 30}, "51.96 30.00"},
	{polar{100, 120}, "-50.00 86.60"},
}

func TestSolver(t *testing.T) {
	// Init solver and channels
	questions := make(chan polar)
	answers := createSolver(questions)
	// Run tests
	for _, tt := range solvertests {
		questions <- tt.in
		c := <-answers
		s := fmt.Sprintf("%.02f %.02f", c.x, c.y)
		if s != tt.out {
			t.Errorf("%v => {%s}, expected: {%s}", tt.in, s, tt.out)
		}
	}
}
