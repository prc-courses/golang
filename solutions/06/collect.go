package main

import "fmt"

var Days = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Sunday",
}

func Arr15() (arr [15]int) {
	for i := 0; i < 15; i++ {
		arr[i] = i
	}
	return
}

func Fib15() (fib [15]int) {
	fib[0], fib[1] = 0, 1

	for i := 2; i < 15; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return
}

func Fibs(n int) (fib []int) {
	fib = make([]int, n)

	if n == 1 {
		fib[0] = 1
		return
	}

	fib[0], fib[1] = 0, 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return
}

func FindDay(day int) string {
	if day, ok := Days[day]; ok {
		return day
	}
	return "Wrong Key"
}

func main() {
	fmt.Println(Arr15())
	fmt.Println(Fib15())
	fmt.Println(Fibs(15))
	fmt.Println(FindDay(4))
	fmt.Println(FindDay(10))
}
