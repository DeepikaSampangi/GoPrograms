package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Guess a random number between 1 and 100 :")
	for i := 0; i <= 10; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("please provide your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil{
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil{
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops! Your guess is low")
		} else if guess > target {
			fmt.Println("Oops! Your guess is high")
		} else {
			fmt.Println("Wow, Congratulations, you got it right !")
			break
		}
	}
	fmt.Println(target)
}
