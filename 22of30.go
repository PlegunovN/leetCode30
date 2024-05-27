package main

import "fmt"

//Write code that enhances all arrays such that you can call the array.last() method on any array and it will return the last element. If there are no elements in the array, it should return -1.

func main() {
	x := make([]int, 0, 0)
	x = append(x, 6, 7, 8, 9, 4)
	if cap(x) == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(x[len(x)-1])
	}
}
