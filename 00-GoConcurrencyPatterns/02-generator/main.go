package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generators are function that returns a channel
func main() {
	c := boring("boring!")
	joe := boring("Joe")
	ann := boring("Ann")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
		fmt.Printf("%q\n", <-joe) // joe will not execute until <-c has received
		fmt.Printf("%q\n", <-ann) // ann will not execute until <-joe has received

	}
	fmt.Println("You're boring, I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		}
	}()
	return c
}
