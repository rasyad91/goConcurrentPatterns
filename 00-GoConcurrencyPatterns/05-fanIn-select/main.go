package main

import (
	"fmt"
	"math/rand"
	"time"
)

// use fan-in function to let whoesover is ready to talk
func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	c := fanIn(joe, ann)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // now will print, on whoever is ready

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

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

// func fanIn(inputs ...<-chan string) <-chan string {
// 	c := make(chan string)
// 	for _, v := range inputs {
// 		go func(v <-chan string) {
// 			for {
// 				c <- <-v
// 			}
// 		}(v)
// 	}
// 	return c
// }
