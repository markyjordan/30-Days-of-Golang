package main

import "fmt"

func concatStrings(str1, str2 string) string {
	return str1 + " " + str2
}

func main() {
	str1 := "Hello,"
	str2 := "World"
	result := concatStrings(str1, str2)
	fmt.Println(result)
}
