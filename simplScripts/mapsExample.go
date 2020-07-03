package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps Example")

	map_ex := make(map[string]int)

	map_ex["a"] = 1
	map_ex["b"] = 2
	fmt.Println(map_ex)

	fmt.Println("Value at a is", map_ex["a"])
	fmt.Println("Key at loc 1 is", map_ex["1"])
	fmt.Println("Total length is", len(map_ex))

	delete(map_ex, "b")
	fmt.Println(map_ex)

}
