package main

import (
	"fmt"
)

func foodPairs() (string, string) {
	return "burger", "frenchFries"
}
func main() {
	fmt.Println(" Multiple values returned by a method")
	var1, var2 := foodPairs()
	fmt.Println("Variety 1 :", var1, "Variety 2 : ", var2)
}
