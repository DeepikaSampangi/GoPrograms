package main

import (
	"fmt"
)

func calculatePaint(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("Input cannot be negative value", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("Input cannot be negative value", width)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	amount, err := calculatePaint(4.2, -3.5)
	fmt.Println(err)
	fmt.Printf("%.2f liters needed\n", amount)
}
