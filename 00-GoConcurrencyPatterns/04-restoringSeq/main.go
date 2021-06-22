package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

// use fan-in function to let whoesover is ready to talk

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	c := fanIn(joe, ann)
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring, I'm leaving.")
}

func boring(msg string) <-chan Message {
	waitForIt := make(chan bool)

	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			c <- Message{
				str:  fmt.Sprintf("%s: %d", msg, i),
				wait: waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Microsecond)
			<-waitForIt
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
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
