package main

import "fmt"

func sliceString(str string, start, end int) string {
	return str[start:end]
}

func main() {
	str := "hello world"
	result := sliceString(str, 0, 5)
	fmt.Println(result)
}
