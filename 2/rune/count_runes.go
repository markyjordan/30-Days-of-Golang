package main

import "fmt"

func countRunes(arr []rune, r rune) int {
	count := 0
	for _, v := range arr {
		if v == r {
			count++
		}
	}
	return count
}

func main() {
	arr := []rune{'a', 'b', 'c', 'd', 'e', 'd', 'c', 'b', 'a'}
	r := rune('c')
	count := countRunes(arr, r)
	fmt.Println("Count of rune", r, "in array:", count)
}
