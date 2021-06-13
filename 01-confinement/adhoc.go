// confinement through a convention, can be from community, group, codebase
// dificult to achieve, unless have tools to perfom static analysis on code on each commit
package main

import "fmt"

func adhoc() {
	data := make([]int, 4)

	loopData := func(dataChan chan<- int) {
		defer close(dataChan)
		for i := range data {
			dataChan <- data[i]
		}
	}

	dataChan := make(chan int)
	go loopData(dataChan)

	for num := range dataChan {
		fmt.Println(num)
	}

}
