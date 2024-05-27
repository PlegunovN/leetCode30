package main

//ok
//Write a function argumentsLength that returns the count of arguments passed to it

import (
	"fmt"
)

func main() {
	var x []string

	for true {
		var y string
		fmt.Scan(&y)
		if y == "exit" {
			break
		}
		x = append(x, y)
	}

	z := argumentsLength(x)
	fmt.Println(z)
}
func argumentsLength(x []string) int {
	y := len(x)
	return y
}
