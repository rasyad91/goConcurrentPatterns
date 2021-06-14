package main

import (
	"fmt"
	"time"
)

func naive() {
	var or func(channels ...<-chan interface{}) <-chan interface{} // takes in variadic slice of and returns a single chan
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0: // since we have recursion, setting up base case
			return nil
		case 1: // if slice has 1 element, return that element
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() { // main body of function, create goroutine so function can wait without blocking
			defer close(orDone)

			switch len(channels) {
			case 2: // because of recursion, every recursive call to "or" will have at least 2 chans.
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...): // recursively creates "or" channel from all the channels in the slice after the 3rd index.
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))

}
