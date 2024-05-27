package main

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

//var z, x, y int
//
//func main() {
//
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			if x == 0 {
//				x = 1
//				x = fn1()
//			}
//		}(i)
//	}
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			if y == 0 {
//				y = 1
//				y = fn2()
//			}
//		}(i)
//	}
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			if z == 0 {
//				z = 1
//				z = fn3()
//			}
//		}(i)
//	}
//
//	time.Sleep(time.Second * 10)
//}
//
//func fn1() int {
//	fmt.Println("fn1")
//	time.Sleep(time.Second * 1)
//	return 0
//}
//
//func fn2() int {
//	fmt.Println("fn2")
//	time.Sleep(time.Second * 2)
//	return 0
//}
//func fn3() int {
//	fmt.Println("fn3")
//	time.Sleep(time.Second * 3)
//	return 0
//}

func fn1() {
	time.Sleep(time.Second * 1)
	fmt.Println("fn1", time.Now())
}

func fn2() {
	time.Sleep(time.Second * 2)
	fmt.Println("fn2", time.Now())
}

func fn3() {
	time.Sleep(time.Second * 3)
	fmt.Println("fn3", time.Now())
}

// Функция дребезга
func debounce(fn func()) {
	var lastRan time.Time
	// Булева переменная для отслеживания состояния функции
	var isRunning bool = false

	for i := 0; i < 10; i++ {
		go func() {
			if !isRunning {
				lastRan = time.Now()
				isRunning = true
				fn()
			} else {
				// Если функция уже запущена, проверяем, прошло ли достаточно времени с последнего вызова
				now := time.Now().Sub(lastRan)
				if now > time.Second {
					lastRan = time.Now()
					isRunning = false
					fn()
				}
			}
		}()
	}
}

func main() {
	debounce(fn1)
	debounce(fn2)
	debounce(fn3)
	time.Sleep(time.Second * 1)
	debounce(fn1)
	debounce(fn2)
	debounce(fn3)
	time.Sleep(time.Second * 1)
	debounce(fn1)
	debounce(fn2)
	debounce(fn3)
	time.Sleep(time.Second * 1)
	debounce(fn1)
	debounce(fn2)
	debounce(fn3)
	time.Sleep(time.Second * 1)
	debounce(fn1)
	debounce(fn2)
	debounce(fn3)

	time.Sleep(time.Second * 10)
}
