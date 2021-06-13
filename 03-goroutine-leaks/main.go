package main

import "fmt"

func main() {
	fmt.Println("======== start of leak example ========")
	leak()
	fmt.Println("======== end of leak example ========") // dowork not exited
	fmt.Println("======== start of non-leak example ========")
	nonleak()
	fmt.Println("======== end of non-leak example ========") // do work exited

	fmt.Println("======== start of leak 2 example ========")
	leak2()
	fmt.Println("======== end of leak 2 example ========") // dowork not exited

	fmt.Println("======== start of non-leak 2 example ========")
	nonleak2()
	fmt.Println("======== end of non-leak 2 example ========") // do work exited

}

// output
// ======== start of leak example ========
// Done.
// ======== end of leak example ========
// ======== start of non-leak example ========
// Cancelling doWork goroutine...
// Exit doWork
// Done.
// ======== end of non-leak example ========

// defer statement never gets run
// ======== start of leak 2 example ========
// 3 random ints:
// 1: 5577006791947779410
// 2: 8674665223082153551
// 3: 6129484611666145821
// ======== end of leak 2 example ========
