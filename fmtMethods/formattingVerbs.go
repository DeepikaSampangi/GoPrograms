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

	fmt.Println("** Formatting values width **")
	fmt.Printf("%%15s : %15s\n", "Hello")
	fmt.Printf("%%4d : %4d\n", 20)
	fmt.Printf("%%7.3f : %7.3f \n", 15.6754)
	fmt.Printf("%%1.3f : %1.3f \n", 15.6754)
	fmt.Printf("%%7.2f : %7.2f \n", 15.6754)
	fmt.Printf("%%7.1f : %7.1f \n", 15.6754)
	fmt.Printf("%%.1f : %.1f \n", 15.6754)
	fmt.Printf("%%.2f : %.2f \n", 15.6754)
}
