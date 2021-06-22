package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(20); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

func boring(msg string, q chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				fmt.Println("case c<-")
			case x := <-q:
				fmt.Println("quit called:", x)
				return
			}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Microsecond)
		}
	}()
	return c
}
