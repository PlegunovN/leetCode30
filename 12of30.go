package main

//ok

import (
	"fmt"
	"sync"
	"time"
)

// Given two promises promise1 and promise2, return a new promise.
//
//	promise1 and promise2 will both resolve with a number. The returned promise should resolve with the sum of the two numbers.
var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	promise1 := 5
	promise2 := 10
	res1 := fn1(promise1)
	res2 := fn2(promise2)
	wg.Wait()
	result := res1 + res2
	fmt.Println(result)
}

func fn1(promise int) int {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	return promise
}
func fn2(promise int) int {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	return promise
}
