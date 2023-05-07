package main

import (
	"fmt"
	"strings"
)

func containsSubstring(str, substr string) bool {
	return strings.Contains(str, substr)
}

func main() {
	str := "hello world"
	substr := "world"
	result := containsSubstring(str, substr)
	fmt.Println(result)
}
