package main

import (
	"errors"
	"fmt"
	// "log"
)

func main() {
	err1 := errors.New("Input values cant be negative")
	fmt.Println(err1.Error())

	// err2 := errors.New("Input values cant be negative")
	// fmt.Println(err2)
	// log.Fatal(err2)

	fmt.Printf("Area of 4.2 by 3.0 %.2f\n", calucalteArea(4.2, 3.0))

	fmt.Printf("Area of 1.2 by 0.5 %.2f\n", calucalteArea(1.2, 0.5))

	fmt.Printf("Area of 4.2 by 3.0 %.2f\n", calucalteArea(5.2, 3.5))

	fmt.Printf("Area of 4.2 by 3.0 %.2f\n", calucalteArea(5.0, 3.3))

}

func calucalteArea(width float64, height float64) float64 {
	return width * height
}
