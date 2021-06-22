package main

import (
	"fmt"
	"math/rand"
	"time"
)

// example shows how channel in go provides a connection between 2 goroutines
// connects the main and boring goroutines so they can communicate
func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 6; i++ {
		// <-c will wait for a value to be sent to channel c
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring, I'm leaving.")
}

func boring(msg string, c chan string) {

	for i := 0; ; i++ {
		// sending value to channel c
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
	}
}
