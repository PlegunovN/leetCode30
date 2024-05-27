package main

//ok
import (
	"fmt"
	"strings"
)

//Given an integer n, return a counter function. This counter function initially returns n and then returns 1 more than the previous value every subsequent time it is called (n, n + 1, n + 2, etc).

func main() {
	var n int
	var c string
	fmt.Scan(&n)
	fmt.Scan(&c)

	words := strings.Split(c, ",")
	for range words {
		counter(n)
		n++
	}

}

func counter(n int) {
	fmt.Println(n)
}
