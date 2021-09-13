package main

import "fmt"

func SumProdDiff(i, j int) (s int, p int, d int) { // named returns
	// Your code here
	return
}

func SumInts(list ...int) (sum int) {
	// Your code here
	return
}

func Factorial(n uint64) uint64 {
	// Your code here
	return 0
}

func Migrate(src []int, fn func(int) int) []int {
	res := make([]int, len(src))
	// Your code here
	return res
}

func square(val int) int {
	return val * val
}

func main() {
	sum, prod, diff := SumProdDiff(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
	// multiple args
	fmt.Println("Sum all:", SumInts(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	// factorial
	fmt.Println("Factorial of 10:", Factorial(10))
	// migrate
	fmt.Println("[1 2 3] squared:", Migrate([]int{1, 2, 3}, square))
}
