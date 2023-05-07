package main

import "fmt"

func countBytes(arr []byte, b byte) int {
	count := 0
	for _, v := range arr {
		if v == b {
			count++
		}
	}
	return count
}

func main() {
	arr := []byte{1, 2, 3, 4, 5, 4, 3, 2, 1}
	b := byte(4)
	count := countBytes(arr, b)
	fmt.Println("Count of byte", b, "in array:", count)
}
