package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	num := 17
	if isPrime(num) {
		fmt.Println(num, "is prime")
	} else {
		fmt.Println(num, "is not prime")
	}
}
