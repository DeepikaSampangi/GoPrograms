package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println("Declaring a variable")
	var x int
	var a,b string
	 x = 1
	 a,b = "a", "b"
	 fmt.Println("value of x is",x,"Type of x is",reflect.TypeOf(x))
	 fmt.Println("value of a is",a)
	 fmt.Println("value of b is",b)


}