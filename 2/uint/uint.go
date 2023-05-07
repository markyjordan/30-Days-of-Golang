package main

import "fmt"

func main() {
	var ui uint = 10
	var ui8 uint8 = 20
	var ui16 uint16 = 30
	var ui32 uint32 = 40
	var ui64 uint64 = 50

	fmt.Println("uint:", ui)
	fmt.Println("uint8:", ui8)
	fmt.Println("uint16:", ui16)
	fmt.Println("uint32:", ui32)
	fmt.Println("uint64:", ui64)

	// Trying to assign ui64 to ui16 results in a compilation error
	// ui16 = ui64 // Error: cannot use ui64 (type uint64) as type uint16 in assignment

	// Trying to assign ui16 to ui8 results in a compilation error
	// ui8 = ui16 // Error: cannot use ui16 (type uint16) as type uint8 in assignment

	// Trying to assign ui8 to ui results in a compilation error
	// ui = ui8 // Error: cannot use ui8 (type uint8) as type uint in assignment

	// Trying to perform arithmetic operations between different unsigned integer types results in a compilation error
	// sum := ui + ui64 // Error: invalid operation: ui + ui64 (mismatched types uint and uint64)

	// However, it is possible to perform arithmetic operations between unsigned integer types of the same size
	sum := uint64(ui) + ui64
	fmt.Println("Sum:", sum)
}
