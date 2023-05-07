package main

import "fmt"

func isVowel(r rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}

func main() {
	r1 := rune('a')
	r2 := rune('b')
	fmt.Println(r1, "is a vowel?", isVowel(r1))
	fmt.Println(r2, "is a vowel?", isVowel(r2))
}
