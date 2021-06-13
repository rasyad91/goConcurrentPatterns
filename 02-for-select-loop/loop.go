/*
for { // either loop infinitely or range over something
	select {
		// Do some work with channels
	}
}


*/

package main

import "fmt"

func main() {

	letters := []string{"a", "b", "c"}
	fmt.Printf("letters slice: %#v\n", letters)
	stringChan := make(chan string)
	done := make(chan bool)

	//Sending iteration variables out on a channel
	for _, s := range letters {
		select {
		case <-done:
			return
		case stringChan <- s:
		}
	}

	// looping infitely waiting to be stopped
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("in default")
		}
		// Do non-preemptable work
	}

}
