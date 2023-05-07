package main

import "fmt"

func strToRunes(str string) []rune {
	return []rune(str)
}

func runesToStr(arr []rune) string {
	return string(arr)
}

func main() {
	str := "hello world"
	arr := strToRunes(str)
	str2 := runesToStr(arr)
	fmt.Println("String:", str)
	fmt.Println("Rune array:", arr)
	fmt.Println("String from rune array:", str2)
}
