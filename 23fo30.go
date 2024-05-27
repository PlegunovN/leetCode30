package main

import "fmt"

//Write code that enhances all arrays such that you can call the array.groupBy(fn) method on any array and it will return a grouped version of the array.
//
//A grouped array is an object where each key is the output of fn(arr[i]) and each value is an array containing all items in the original array with that key.
//
//The provided callback fn will accept an item in the array and return a string key.
//
//The order of each value list should be the order the items appear in the array. Any order of keys is acceptable.
//
//Please solve it without lodash's _.groupBy function.

type Array struct {
	values []interface{}
}

func (a *Array) groupBy(fn func(interface{}) string) map[string][]interface{} {
	result := make(map[string][]interface{})
	for _, value := range a.values {
		key := fn(value)
		if _, ok := result[key]; !ok {
			result[key] = []interface{}{}
		}
		result[key] = append(result[key], value)
	}
	return result
}

func main() {
	a := Array{
		values: []interface{}{1, 2, 3, 8, 4, 5, 9},
	}

	b := a.groupBy(func(v interface{}) string {
		return fmt.Sprintf("%d", v.(int))
	})

	fmt.Println(b)
}
