package main

import "fmt"

//Given an object or an array, return if it is empty.
//
//An empty object contains no key-value pairs.
//An empty array contains no elements.
//
//You may assume the object or array is the output of JSON.parse.

//	func isEmpty(obj interface{}) bool {
//		switch obj := obj.(type) {
//		case map[string]interface{}:
//			return true
//		default:
//			if reflect.
//			//if reflect.TypeOf(obj).Kind() == reflect.Array {
//				return true
//			} else {
//				return false
//			}
//		}
//	}
func main() {
	var obj1 map[string]int // Пустой объект
	arr1 := []int{}         // Пустой массив

	fmt.Println("Объект1:", mapEmpty(obj1))
	fmt.Println("Массив1:", arrEmpty(arr1))

	obj2 := map[string]int{
		"key1": 1,
		"key2": 2,
	} // Заполненный объект
	arr2 := []int{1, 2, 3} // Заполненный массив

	fmt.Println("Объект2:", mapEmpty(obj2))
	fmt.Println("Массив2:", arrEmpty(arr2))
}

func arrEmpty(arr []int) bool {
	if len(arr) == 0 {
		return true
	} else {
		return false
	}

}

func mapEmpty(obj map[string]int) bool {
	if obj == nil {
		return true
	} else {
		return false
	}
}
