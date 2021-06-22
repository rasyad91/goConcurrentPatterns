package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now()
	results := Googlev1("golang")
	elasped := time.Since(start)
	fmt.Println(results)
	fmt.Println(elasped)

	start2 := time.Now()
	results2 := Googlev2("golang")
	elasped2 := time.Since(start2)
	fmt.Println(results2)
	fmt.Println(elasped2)

	start3 := time.Now()
	results3 := Googlev3("golang")
	elasped3 := time.Since(start3)
	fmt.Println(results3)
	fmt.Println(elasped3)
}
