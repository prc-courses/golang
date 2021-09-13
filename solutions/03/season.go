package main

import "fmt"

func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "Winter" // Dec, Jan and Feb have winter
	case 3, 4, 5:
		return "Spring" // March, Apr and May have spring
	case 6, 7, 8:
		return "Summer" // June, July and Aug have summer
	case 9, 10, 11:
		return "Autumn" // Sept, Oct and Nov have autumn

	default:
		return "Season unknown" // value outside [1,12], then season is unkown
	}
}

func main() {
	for i := 1; i <= 12; i += 3 {
		fmt.Println(Season(i))
	}
}
