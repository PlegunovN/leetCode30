package main

//ok

import "fmt"

//Given an integer array arr and a mapping function fn, return a new array with a transformation applied to each element.
//
//The returned array should be created such that returnedArray[i] = fn(arr[i], i).

func main() {
	var arr []int
	output := fn(arr)

	for true {
		var y int
		fmt.Scan(&y)
		if y == -0 {
			break
		}
		arr = append(arr, y)
	}

	fmt.Println(output)
}

func fn(arr []int) []int {
	var output []int
	for i := 0; i < len(arr); i++ {
		output = append(output, (arr[i] + 1))
	}

	return output
}
