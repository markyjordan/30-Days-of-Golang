package main

import "fmt"

func main() {
	var i int = 10
	var i8 int8 = 20
	var i16 int16 = 30
	var i32 int32 = 40
	var i64 int64 = 50

	fmt.Println("int:", i)
	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)

	// Trying to assign i64 to i16 results in a compilation error
	// i16 = i64 // Error: cannot use i64 (type int64) as type int16 in assignment

	// Trying to assign i16 to i8 results in a compilation error
	// i8 = i16 // Error: cannot use i16 (type int16) as type int8 in assignment

	// Trying to assign i8 to i results in a compilation error
	// i = i8 // Error: cannot use i8 (type int8) as type int in assignment

	// Trying to perform arithmetic operations between different integer types results in a compilation error
	// sum := i + i64 // Error: invalid operation: i + i64 (mismatched types int and int64)

	// However, it is possible to perform arithmetic operations between integer types of the same size
	sum := int64(i) + i64
	fmt.Println("Sum:", sum)
}
