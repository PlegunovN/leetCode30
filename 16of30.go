package main

//Given an asynchronous function fn and a time t in milliseconds, return a new time limited version of the input function. fn takes arguments provided to the time limited function.
//
//The time limited function should follow these rules:
//
//If the fn completes within the time limit of t milliseconds, the time limited function should resolve with the result.
//If the execution of the fn exceeds the time limit, the time limited function should reject with the string "Time Limit Exceeded"
import (
	"fmt"
	"os"
	"time"
)

func main() {
	var t time.Duration
	t = 1
	go fn2(t)
	l := 100
	for i := 0; i < l; i++ {
		fn(i)
		time.Sleep(time.Millisecond * 300)

	}

}
func fn(i int) {
	fmt.Println(i)
}
func fn2(t time.Duration) {
	time.Sleep(time.Second * t)
	fmt.Println("Time Limit Exceeded")
	os.Exit(0)
}
