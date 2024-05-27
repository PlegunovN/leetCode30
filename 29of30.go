package main

import (
	"fmt"
	"strings"
)

//Create a class ArrayWrapper that accepts an array of integers in its constructor. This class should have two features:
//
//When two instances of this class are added together with the + operator, the resulting value is the sum of all the elements in both arrays.
//When the String() function is called on the instance, it will return a comma separated string surrounded by brackets. For example, [1,2,3].

type ArrayWrapper struct {
	array []int
}

func (aw *ArrayWrapper) Add(other ArrayWrapper) int {
	var result int
	for i := 0; i < len(aw.array); i++ {
		result = result + aw.array[i]
	}
	for i := 0; i < len(other.array); i++ {
		result = result + other.array[i]
	}

	return result
}

func (aw ArrayWrapper) String() string {
	var result strings.Builder
	result.WriteString("[")
	for i, v := range aw.array {
		if i > 0 {
			result.WriteByte(',')
		}
		result.WriteString(fmt.Sprintf("%d", v))
	}
	result.WriteString("]")

	return result.String()
}

func main() {
	aw1 := ArrayWrapper{[]int{1, 2, 3}}
	aw2 := ArrayWrapper{[]int{4, 5, 6}}

	fmt.Println(aw1.Add(aw2))
	fmt.Println(aw1.String())

}
