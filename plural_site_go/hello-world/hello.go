package main

import (
	"fmt"
)

func main() {
	var (
		test string = "2"
		test2 string = "2"
	)
	fmt.Println("Original", test)
	fmt.Println("Returned", testing(test))
	fmt.Println("Original", test2)
	fmt.Println("Returned", testing(test2))
}

func testing(something string) string {
	something = "changed"
	return something
}