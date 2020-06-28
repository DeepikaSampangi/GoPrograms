package main

import (
	"fmt"
)

func main() {
	fmt.Printf("A float Number %f \n", 1.23412)
	fmt.Printf("An Integer %d \n", 32)
	fmt.Printf("A String %s \n", "Steve")
	fmt.Printf("A Boolean value right : %t, wrong : %t \n", true, "hello")
	fmt.Printf("Any Value %v, %v, %v \n", 4.9, "\n", false)
	fmt.Printf("Any Vaule decalred %#v, %#v, %#v \n", 4.9, "\t", false)
	fmt.Printf("Types of the values 1.2 : %T, tab space : %T, Bool value : %T \n", 1.2, "\t", true)
	fmt.Printf("Percent sign : %% \n")
}
