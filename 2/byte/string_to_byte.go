package main

import "fmt"

func strToBytes(str string) []byte {
	return []byte(str)
}

func bytesToStr(arr []byte) string {
	return string(arr)
}

func main() {
	str := "hello world"
	arr := strToBytes(str)
	str2 := bytesToStr(arr)
	fmt.Println("String:", str)
	fmt.Println("Byte array:", arr)
	fmt.Println("String from byte array:", str2)
}
