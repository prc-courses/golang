package main

import "fmt"

// aliasing type
type Celsius float32
type Fahrenheit float32

// Function to convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit = Fahrenheit(t*9/5) + 32
	return temp
}

// Main
func main() {
	fmt.Println("Temperature 100ºC in Farenheit:", toFahrenheit(100))
}
