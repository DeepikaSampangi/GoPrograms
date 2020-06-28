package main

import (
	"fmt"
)

func main() {
	val := 8
	fmt.Println("Values is", val)
	doubleTheValue(val)
	fmt.Println("Without pointer multiply", val)

	doubleTheValuePointer(&val)
	fmt.Println("With pointer multiply", val)
}

func doubleTheValue(val int) int {
	return val * 2
}

func doubleTheValuePointer(val *int) {
	*val *= 2
}
