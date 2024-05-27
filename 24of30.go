package main

import (
	"fmt"
	"sort"
)

// Given an array arr and a function fn, return a sorted array sortedArr. You can assume fn only returns numbers and those numbers determine the sort order of sortedArr.
// sortedArray must be sorted in ascending order by fn output.
//
// You may assume that fn will never duplicate numbers for a given array.
func main() {
	arr := []int{3, 6, 50, 2, 7}
	sort.Ints(arr)
	fmt.Println(arr)

}
