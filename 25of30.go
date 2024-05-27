package main

import "fmt"

//Given two arrays arr1 and arr2, return a new array joinedArray. All the objects in each of the two inputs arrays will contain an id field that has an integer value.
//
//joinedArray is an array formed by merging arr1 and arr2 based on their id key. The length of joinedArray should be the length of unique values of id. The returned array should be sorted in ascending order based on the id key.
//
//If a given id exists in one array but not the other, the single object with that id should be included in the result array without modification.
//
//If two objects share an id, their properties should be merged into a single object:
//
//If a key only exists in one object, that single key-value pair should be included in the object.
//If a key is included in both objects, the value in the object from arr2 should override the value from arr1.

func main() {

	arr1 := []map[string]int{
		{"id": 1},
		{"x": 2},
		{"y": 3},
	}
	arr2 := []map[string]int{
		{"id": 2},
		{"x": 4},
		{"y": 5},
	}
	startArray := [][]map[string]int{{arr1[0], arr1[1], arr1[2]}, {arr2[0], arr2[1], arr2[2]}}

	arr3 := [3]map[string]int{
		{"id": 3},
		{"x": 6},
		{"y": 7},
	}
	newArray := [][]map[string]int{{arr3[0], arr3[1], arr3[2]}}
	finalArray := make([][]map[string]int, len(startArray)+len(newArray))

	for i := 0; i < len(startArray); i++ {
		finalArray[i] = startArray[i]
	}

	for j := 0; j < len(newArray); j++ {
		finalArray[len(startArray)+j] = newArray[j]
	}

	fmt.Println(finalArray)
}
