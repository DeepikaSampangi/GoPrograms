package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go Slices Operations")

	s := make([]string, 3)
	fmt.Println("Elements in set are : ", s)

	s[0] = "Hello"
	s[1] = "Hey"
	s[2] = "Hai"
	fmt.Println("Elements in set are : ", s)
	fmt.Println("Length of set", len(s))
}
