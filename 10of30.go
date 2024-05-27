package main

// уточнить
import "fmt"

//Given a function fn, return a new function that is identical to the original function except that it ensures fn is called at most once.
//
//The first time the returned function is called, it should return the same result as fn.
//Every subsequent time it is called, it should return undefined

func main() {
	l := 10
	max := 1
	x := 3
	y := 5
	z := 9
	for i := 0; i < l; i++ {
		if i < max {
			firstCall := fn(x, y, z)
			fmt.Println(firstCall)
		} else {
			fmt.Println("undefined, fn was not called")
		}
	}
}

func fn(x, y, z int) int {
	n := x * y * z
	return n
}
