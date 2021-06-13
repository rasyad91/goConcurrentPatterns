// expose only the correct data and concurrency primitives for multiple concurrent processes to use.
// makes it impossible to do the wrong thing

package main

import (
	"bytes"
	"fmt"
	"sync"
)

func lexical1() {
	producer := func() <-chan int {
		// 1) instantiate channel within lexical scope of producer.
		//confines write aspect of this channel in this function to prevent other goroutines from writing to it
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 6; i++ {
				results <- i
			}
		}()
		return results
	}

	// 3) receive a read-only copy of an int channel.
	// by declaring that the only usage we require is read access, it confines the usage of the channel
	// within the consume function to only reads
	consumer := func(results <-chan int) {
		for v := range results {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	// 2) receive ONLY the read aspect of the channel and pass it to consumer
	// which can only read from it. Confines main goroutine to a read-only view
	results := producer()
	consumer(results)
}

func lexical2() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()

}
