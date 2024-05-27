package main

import (
	"fmt"
	"time"
)

// Given a positive integer millis, write an asynchronous function that sleeps for millis milliseconds. It can resolve any value.
func main() {
	input := 100
	time.Sleep(time.Millisecond * 100)
	fmt.Println(input)
}
