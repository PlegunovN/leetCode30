package main

import (
	"fmt"
	"os"
	"time"
)

// Given a function fn, an array of arguments args, and an interval time t, return a cancel function cancelFn.
//
// After a delay of cancelTimeMs, the returned cancel function cancelFn will be invoked.
//
// setTimeout(cancelFn, cancelTimeMs)
//
// The function fn should be called with args immediately and then called again every t milliseconds until cancelFn is called at cancelTimeMs ms.

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
	fmt.Println("timer off")
	os.Exit(0)
}
