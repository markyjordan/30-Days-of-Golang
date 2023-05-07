package main

import "fmt"

func addComplex64(c1 complex64, c2 complex64) complex64 {
	return c1 + c2
}

func addComplex128(c1 complex128, c2 complex128) complex128 {
	return c1 + c2
}

func main() {
	c1_64 := complex(float32(2), float32(3))
	c2_64 := complex(float32(4), float32(5))
	c1_128 := complex(float64(2), float64(3))
	c2_128 := complex(float64(4), float64(5))
	sum64 := addComplex64(c1_64, c2_64)
	sum128 := addComplex128(c1_128, c2_128)
	fmt.Println("Sum of complex numbers", c1_64, "and", c2_64, "(complex64) is", sum64)
	fmt.Println("Sum of complex numbers", c1_128, "and", c2_128, "(complex128) is", sum128)
}
