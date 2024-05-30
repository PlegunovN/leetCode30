package main

//ok
import (
	"fmt"
	"time"
)

//Given a function fn and a time in milliseconds t, return a debounced version of that function.
//
//A debounced function is a function whose execution is delayed by t milliseconds and whose execution is cancelled if it is called again within that window of time. The debounced function should also receive the passed parameters.
//
//For example, let's say t = 50ms, and the function was called at 30ms, 60ms, and 100ms.
//
//The first 2 function calls would be cancelled, and the 3rd function call would be executed at 150ms.
//
//If instead t = 35ms, The 1st call would be cancelled, the 2nd would be executed at 95ms, and the 3rd would be executed at 135ms.

func main() {

	var x = 0
	var y = 0
	var z = 0

	go func() {
		for i := 0; i < 10; {
			if x == 0 {
				x = 1
				go func(i int) {
					x = fn1(x)
				}(i)
				i++
			}
		}
	}()

	go func() {
		for i := 0; i < 10; {
			if y == 0 {
				y = 1
				go func(i int) {
					y = fn2(y)
				}(i)
				i++
			}
		}
	}()
	go func() {
		for i := 0; i < 10; {
			if z == 0 {
				z = 1
				go func(i int) {
					z = fn3(z)
				}(i)
				i++
			}
		}
	}()

	time.Sleep(20 * time.Second)
}

func fn1(x int) int {
	fmt.Println("fn1", time.Now())
	time.Sleep(time.Second * 1)
	return 0
}

func fn2(x int) int {
	fmt.Println("fn2", time.Now())
	time.Sleep(time.Second * 2)
	return 0
}
func fn3(x int) int {
	fmt.Println("fn3", time.Now())
	time.Sleep(time.Second * 3)
	return 0
}
