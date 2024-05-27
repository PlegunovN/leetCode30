package main

//ok

import (
	"fmt"
	"log"
	"strings"
)

//Write a function createCounter. It should accept an initial integer init. It should return an object with three functions.
//
//The three functions are:
//
//increment() increases the current value by 1 and then returns it.
//decrement() reduces the current value by 1 and then returns it.
//reset() sets the current value to init and then returns it.

func main() {
	x := 0
	fmt.Scan(&x)
	y := x
	calls := ""
	fmt.Scan(&calls)
	var out []int
	var words []string

	calls = strings.ReplaceAll(calls, `"`, "")
	calls = strings.ReplaceAll(calls, `[`, "")
	calls = strings.ReplaceAll(calls, `]`, "")

	words = strings.Split(calls, ",")

	for i := 0; i < len(words); i++ {
		if (strings.Compare(words[i], "increment")) == 0 {
			x = increment(x)
		} else if (strings.Compare(words[i], "decrement")) == 0 {
			x = decrement(x)
		} else if (strings.Compare(words[i], "reset")) == 0 {
			x = reset(x, y)
		} else {
			log.Fatal("err")
		}
		out = append(out, x)
	}
	fmt.Println(out)
}

func increment(x int) int {
	x++
	return x
}

func decrement(x int) int {
	x--
	return x
}

func reset(x, y int) int {
	x = y
	return x
}
