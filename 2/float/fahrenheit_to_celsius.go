package main

import "fmt"

func fahrenheitToCelsius32(fahrenheit float32) float32 {
	return (fahrenheit - 32) * 5 / 9
}

func fahrenheitToCelsius64(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	temp32 := float32(72.0)
	temp64 := float64(72.0)
	celsius32 := fahrenheitToCelsius32(temp32)
	celsius64 := fahrenheitToCelsius64(temp64)
	fmt.Println(temp32, "degrees Fahrenheit (float32) is", celsius32, "degrees Celsius")
	fmt.Println(temp64, "degrees Fahrenheit (float64) is", celsius64, "degrees Celsius")
}
