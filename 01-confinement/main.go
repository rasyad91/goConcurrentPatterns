// Confinement - ensuring information is only ever available from one concurrent process
// Which will make the concurrent program implicitly safe and no synchronization needed
// why confinement if we have synchronization available?
// 1) Improved performance 2) reduce cognitive load on developers
package main

import "fmt"

func main() {
	fmt.Println("======== start of adhoc example ========")
	adhoc()
	fmt.Println("======== end of adhoc example ========")
	println()

	fmt.Println("======== start of lexical 1 example ========")
	lexical1()
	fmt.Println("======== end of lexical 1 example ========")
	println()

	fmt.Println("======== start of lexical 2 example ========")
	lexical2()
	fmt.Println("======== end of lexical 2 example ========")

}
