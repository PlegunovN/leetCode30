//Given an array of functions [f1, f2, f3, ..., fn], return a new function fn that is the function composition of the array of functions.
//
//The function composition of [f(x), g(x), h(x)] is fn(x) = f(g(h(x))).
//
//The function composition of an empty list of functions is the identity function f(x) = x.
//
//You may assume each function in the array accepts one integer as input and returns one integer as output.

package main

//ok
import "fmt"

type Function func(int) int

func compose(fs []Function) Function {
	if len(fs) == 0 {
		return func(x int) int { return x }
	} else {

		return func(x int) int {
			var result int
			for _, f := range fs {
				result = f(result)

			}
			return result
		}
	}
}
func main() {
	f1 := func(x int) int { return x + 1 }
	f2 := func(x int) int { return x / 2 }
	f3 := func(x int) int { return x * 3 }
	f4 := func(x int) int { return x + 3 }

	fn := compose([]Function{f4, f3, f2, f1})
	fmt.Println(fn(0))
}
