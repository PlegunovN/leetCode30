package main

//ok

import "fmt"

//Given an array arr and a chunk size size, return a chunked array.
//
//A chunked array contains the original elements in arr, but consists of subarrays each of length size. The length of the last subarray may be less than size if arr.length is not evenly divisible by size.
//
//You may assume the array is the output of JSON.parse. In other words, it is valid JSON.
//
//Please solve it without using lodash's _.chunk function.

func main() {
	var arr []int
	var x int
	fmt.Println("x=")
	fmt.Scan(&x)

	var out [][]int
	fmt.Println("arr=")
	for true {
		var y int
		fmt.Scan(&y)
		if y == -0 {
			break
		}
		arr = append(arr, y)
	}

	l := len(arr)/x + len(arr)%x

	for i := 0; i < l; i++ {
		o := fn(arr, x)
		out = append(out, o)
		arr = remove(arr, x)
	}
	fmt.Println(out)
}

func fn(arr []int, x int) []int {
	var o []int
	if len(arr) < x {
		x = len(arr)
	}
	for i := 0; i < x; i++ {
		o = append(o, arr[i])

	}
	return o
}
func remove(arr []int, x int) []int {
	var nArr []int
	for i := 0; i < len(arr); i++ {
		if i >= x {
			nArr = append(nArr, arr[i])
		}
	}
	return nArr
}
