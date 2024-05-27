package main

//ok
import "fmt"

//Given an integer array nums, a reducer function fn, and an initial value init, return the final result obtained by executing the fn function on each element of the array, sequentially, passing in the return value from the calculation on the preceding element.
//
//This result is achieved through the following operations: val = fn(init, nums[0]), val = fn(val, nums[1]), val = fn(val, nums[2]), ... until every element in the array has been processed. The ultimate value of val is then returned.
//
//If the length of the array is 0, the function should return init.

func main() {
	var nums []int
	var init int
	var val int
	for true {
		var y int
		fmt.Scan(&y)
		if y == 0 {
			break
		}
		nums = append(nums, y)
	}
	fmt.Scan(&init)

	if len(nums) == 0 {
		fmt.Println(init)
	} else {
		val = fn(init, nums[0])
		for i := 1; i < len(nums); i++ {
			val = fn(val, nums[i])
		}
		fmt.Println(val)
	}
}

func fn(x, y int) int {
	return x + y
}
