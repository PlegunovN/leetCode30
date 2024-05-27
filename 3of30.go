package main

import "fmt"

//Write a function expect that helps developers test their code. It should take in any value val and return an object with the following two functions.
//
//toBe(val) accepts another value and returns true if the two values === each other. If they are not equal, it should throw an error "Not Equal".
//notToBe(val) accepts another value and returns true if the two values !== each other. If they are equal, it should throw an error "Equal".

type Object struct {
	val  int
	val2 int
}

func main() {

	object1 := Object{2, 3}
	object1.toBe()
	object1.notToBe()
}

func (object Object) toBe() {
	if object.val == object.val2 {
		fmt.Println("each other")
	} else {
		fmt.Println("Not Equal")
	}
}

func (object Object) notToBe() {
	if object.val == object.val2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("each other")
	}
}
