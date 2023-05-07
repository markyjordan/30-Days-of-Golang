package main

import "fmt"

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	str := "racecar"
	if isPalindrome(str) {
		fmt.Println(str, "is a palindrome")
	} else {
		fmt.Println(str, "is not a palindrome")
	}
}
