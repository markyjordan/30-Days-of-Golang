package main

import (
	"fmt"
	"strings"
)

func toUppercase(str string) string {
	return strings.ToUpper(str)
}

func main() {
	str := "hello world"
	result := toUppercase(str)
	fmt.Println(result)
}
