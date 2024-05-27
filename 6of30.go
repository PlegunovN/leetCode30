package main

//ok

//Given an integer array arr and a filtering function fn, return a filtered array filteredArr.
//
//The fn function takes one or two arguments:
//
//arr[i] - number from the arr
//i - index of arr[i]
//
//filteredArr should only contain the elements from the arr for which the expression fn(arr[i], i) evaluates to a truthy value. A truthy value is a value where Boolean(value) returns true.

import "fmt"

func main() {
	var arr []int

	for true {
		var y int
		fmt.Scan(&y)
		if y == -0 {
			break
		}
		arr = append(arr, y)
	}
	var n int
	fmt.Scan(&n)

	var x []int
	x = fn(arr, n)
	fmt.Println(x)
}
func fn(arr []int, n int) []int {
	var x []int
	for i := n; i < len(arr); i++ {
		x = append(x, arr[i])
	}
	return x
}
