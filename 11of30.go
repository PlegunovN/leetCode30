// Given a function fn, return a memoized version of that function.
//
// A memoized function is a function that will never be called twice with the same inputs. Instead it will return a cached value.
//
// You can assume there are 3 possible input functions: sum, fib, and factorial.
//
// sum accepts two integers a and b and returns a + b. Assume that if a value has already been cached for the arguments (b, a) where a != b, it cannot be used for the arguments (a, b). For example, if the arguments are (3, 2) and (2, 3), two separate calls should be made.
// fib accepts a single integer n and returns 1 if n <= 1 or fib(n - 1) + fib(n - 2) otherwise.
// factorial accepts a single integer n and returns 1 if n <= 1 or factorial(n - 1) * n otherwise.
// ok
package main

import "fmt"

type Fun struct {
	a int
	b int
}

func (S *Fun) GetSum(a, b int) int {
	x := a + b
	return x
}

func memoizedGetSum(S *Fun, a, b int) (result int) {
	if result, ok := cache[a*10+b]; ok {
		return result
	} else {
		result = S.GetSum(a, b)
		cache[a*10+b] = result
		return result
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	} else {
		result, ok := cache[n]
		if !ok {
			result = fib(n-1) + fib(n-2)
			cache[n] = result
		}
		return result
	}
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		result, ok := cache[n]
		if !ok {
			result = factorial(n-1) * n
			cache[n] = result
		}
		return result
	}
}

var cache map[int]int

func init() {
	cache = make(map[int]int)
}

func main() {
	fun1 := Fun{3, 5}
	memoizedX := memoizedGetSum(&fun1, fun1.a, fun1.b)
	fmt.Println("Сумма", fun1.a, "и", fun1.b, "=", memoizedX)
	fmt.Println("5-е число Фибоначчи =", fib(5))
	fmt.Println("Факториал 7 =", factorial(7))
}
