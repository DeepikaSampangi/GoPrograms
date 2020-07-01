package main

import (
	"fmt"
)

func addFunc(nums ...int) {
	fmt.Println("Input is", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func main() {
	fmt.Println("Variadic function are func with n number of inputs")

	addFunc(1, 2)
	addFunc(3, 4, 5)

}
