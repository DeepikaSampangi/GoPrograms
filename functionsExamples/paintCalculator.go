package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Area of 4.2 by 3.0 %.2f\n", calucalteArea(4.2, 3.0))

	fmt.Printf("Area of 1.2 by 0.5 %.2f\n", calucalteArea(1.2, 0.5))

	fmt.Printf("Area of 4.2 by 3.0 %.2f\n", calucalteArea(5.2, 3.5))

	fmt.Printf("Area of 4.2 by 3.0 %.2f\n", calucalteArea(5.0, 3.3))

}

func calucalteArea(width float64, height float64) float64 {
	return width * height
}
