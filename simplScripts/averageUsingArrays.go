package main

import (
	"fmt"
)

func main() {
	fmt.Println("Average of an array")

	var myArr [4]float32 = [4]float32{10.2, 24.5, 15.3}
	var sum float32 = 0
	for _, ele := range myArr {
		sum = sum + ele
	}
	total_count := len(myArr)

	fmt.Println("The avearge of ", myArr, "is ", (sum / float32(total_count)))
}
