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

	s = append(s, "World")
	s = append(s, "Universe", "Galaxy")
	fmt.Println("Elements in set after appending : ", s)

	sub_list := s[3:6]
	fmt.Println("Elements in sub list are :", sub_list)

	s2 := []string{"Planets", "Stars", "Comets"}
	fmt.Println("List createad at declaration", s2)

	twoDim := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDim[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDim[i][j] = i + j
		}
	}
	fmt.Println("twoD list", twoDim)
}
