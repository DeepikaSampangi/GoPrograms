package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Today is a ")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Hoorayy ! It's Weekend")
	default:
		fmt.Println("Weekday, Keep the hustle going ! ")
	}

}
