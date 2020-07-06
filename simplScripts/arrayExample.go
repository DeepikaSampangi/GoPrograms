package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Declaring an Array")

	var myArr [5]string

	myArr[0] = "sa"
	myArr[1] = "re"
	myArr[2] = "ga"
	myArr[3] = "ma"
	myArr[4] = "pa"
	fmt.Println(myArr)

	fmt.Println("Array Literals Example")

	var arrLiteral [5]int = [5]int{1, 2, 3}
	fmt.Println("Array Literal", arrLiteral)

	arr := [3]time.Time{time.Unix(1447920000, 0), time.Unix(1447920000, 0), time.Unix(1447920000, 0)}
	fmt.Println(arr)

	fmt.Println("%#v", arr)

	for i := 0; i <= 3; i++ {
		fmt.Println(myArr[i])
	}
	fmt.Println(len(myArr))

}
