package main

//ok
import "fmt"

//Write a function createHelloWorld. It should return a new function that always returns "Hello World".

func main() {
	var x string
	fmt.Scan(&x)
	helloWorld(x)
}

func helloWorld(y string) {
	y = "Hello World"
	fmt.Println(y)
}
