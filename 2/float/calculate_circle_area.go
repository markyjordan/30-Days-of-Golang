package main

import "fmt"

func calculateCircleArea32(radius float32) float32 {
	return 3.14159 * radius * radius
}

func calculateCircleArea64(radius float64) float64 {
	return 3.14159 * radius * radius
}

func main() {
	radius32 := float32(1.0)
	radius64 := float64(1.0)
	area32 := calculateCircleArea32(radius32)
	area64 := calculateCircleArea64(radius64)
	fmt.Println("Area of circle with radius", radius32, "(float32) is", area32)
	fmt.Println("Area of circle with radius", radius64, "(float64) is", area64)
}
