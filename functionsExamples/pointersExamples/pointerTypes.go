package main

import (
	"fmt"
)

func main() {
	int_val := 4
	fmt.Println("Value", int_val)
	int_val_pointer := &int_val
	*int_val_pointer = 8
	fmt.Println("Pointer itself", int_val_pointer)
	fmt.Println("The value stored at the pointer", *int_val_pointer)
	fmt.Println("Value", int_val)
	fmt.Println("------------------------------------------")

	float_val := 4.28
	fmt.Println("Value", float_val)
	float_val_pointer := &float_val
	*float_val_pointer = 8.56
	fmt.Println("Pointer itself", float_val_pointer)
	fmt.Println("The value stored at the pointer", *float_val_pointer)
	fmt.Println("Value", float_val)
	fmt.Println("------------------------------------------")


	bool_val := true
	fmt.Println("Value", bool_val)
	bool_val_pointer := &bool_val
	*bool_val_pointer = false 
	fmt.Println("Pointer itself", bool_val_pointer)
	fmt.Println("The value stored at the pointer", *bool_val_pointer)
	fmt.Println("Value", bool_val)
	fmt.Println("------------------------------------------")


}
