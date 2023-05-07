package main

import "fmt"

func reverseBytes(arr []byte) []byte {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	arr := []byte{1, 2, 3, 4, 5}
	reversed := reverseBytes(arr)
	fmt.Println("Original array:", arr)
	fmt.Println("Reversed array:", reversed)
}
