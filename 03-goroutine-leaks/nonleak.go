package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Establish a signal between parent goroutine and its children
// that allows the parent to signal cancellation to its children.
// By convention, signal is usually a read-only chan named done.
// The parent passes this channel to the child and then closes the chan
// when it wants to cancel the child goroutine.
// receiving case
func nonleak() {

	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} { // 1) pass done chan, as convention, first parameter
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("Exit doWork")
			defer close(terminated)
			for {
				select { // 2) ubiquitous for-select pattern in use. Case statement checking whether done chan has been signaled.
				case s := <-strings:
					//Do something.
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() { // 3) goroutine that will cancel the goroutine spawned in dowork after 1s
		// cancel the operation after 1s
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling doWork goroutine...")
		close(done)
	}()

	<-terminated // 4) join the goroutine spawned from doWork with main goroutine
	fmt.Println("Done.")
}

func nonleak2() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints: ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
}
