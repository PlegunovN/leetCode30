package main

//ok
import "fmt"

// Given a multi-dimensional array arr and a depth n, return a flattened version of that array.
//
// A multi-dimensional array is a recursive data structure that contains integers or other multi-dimensional arrays.
//
// A flattened array is a version of that array with some or all of the sub-arrays removed and replaced with the actual elements in that sub-array.
//
//	This flattening operation should only be done if the current depth of nesting is less than n. The depth of the elements in the first array are considered to be 0.
//
// Please solve it without the built-in Array.flat method.
func main() {

	arr := [2][2][3]int{
		{{1, 2, 3},
			{4, 5, 6}},
		{{1, 2, 3},
			{4, 5, 6}},
	}

	flattened := make([]int, len(arr)*len(arr[0])*len(arr[0][0]))
	i := 0
	for _, val0 := range arr {
		for _, val1 := range val0 {
			for _, val2 := range val1 {
				flattened[i] = val2
				i++
			}
		}
	}

	for _, value := range flattened {
		fmt.Print(value)
	}
}
