package main

//ok
import (
	"fmt"
	"os"
	"time"
)

//Given a function fn, an array of arguments args, and a timeout t in milliseconds, return a cancel function cancelFn.
//
//After a delay of cancelTimeMs, the returned cancel function cancelFn will be invoked.
//
//setTimeout(cancelFn, cancelTimeMs)
//
//Initially, the execution of the function fn should be delayed by t milliseconds.

// If, before the delay of t milliseconds, the function cancelFn is invoked, it should cancel the delayed execution of fn.
// Otherwise, if cancelFn is not invoked within the specified delay t, fn should be executed with the provided args as arguments.
func main() {
	var t time.Duration
	t = 3
	x := 5
	fmt.Println("запущена")
	go fn(x)
	fn2(t)

}
func fn(x int) {
	x = x * x
	time.Sleep(time.Second * 2)
	fmt.Println(x)
}

func fn2(t time.Duration) {
	time.Sleep(time.Second * t)
	fmt.Println("timer off")
	os.Exit(0)
}
