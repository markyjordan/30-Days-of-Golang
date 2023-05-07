package main

import "fmt"

func calculateSquare64(c complex64) complex64 {
	return c * c
}

func calculateSquare128(c complex128) complex128 {
	return c * c
}

func main() {
	c64 := complex(float32(2), float32(3))
	c128 := complex(float64(2), float64(3))
	square64 := calculateSquare64(c64)
	square128 := calculateSquare128(c128)
	fmt.Println("Square of complex number", c64, "(complex64) is", square64)
	fmt.Println("Square of complex number", c128, "(complex128) is", square128)
}
