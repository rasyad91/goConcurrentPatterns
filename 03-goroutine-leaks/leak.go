package main

import (
	"fmt"
	"math/rand"
	"time"
)

// receiving case
func leak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
			fmt.Println("Running in do work")
			time.Sleep(3 * time.Second)

		}()
		return completed
	}

	doWork(nil)
	// perhaps more work is done here...
	fmt.Println("Done.")
}

// sending case
func leak2() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.") // didnt get executed
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints: ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}
