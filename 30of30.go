package main

import (
	"fmt"
	"math"
)

//Design a Calculator class. The class should provide the mathematical operations of addition, subtraction, multiplication, division, and exponentiation.
//	It should also allow consecutive operations to be performed using method chaining. The Calculator class constructor should accept a number which serves as the initial value of result.
//
//Your Calculator class should have the following methods:
//
//add - This method adds the given number value to the result and returns the updated Calculator.
//subtract - This method subtracts the given number value from the result and returns the updated Calculator.
//multiply - This method multiplies the result  by the given number value and returns the updated Calculator.
//divide - This method divides the result by the given number value and returns the updated Calculator.
//	If the passed value is 0, an error "Division by zero is not allowed" should be thrown.
//power - This method raises the result to the power of the given number value and returns the updated Calculator.
//getResult - This method returns the result.
//
//Solutions within 10-5 of the actual result are considered correct.

type Calculator struct {
	result float64
}

func (c *Calculator) Add(a float64) *Calculator {
	c.result += a
	return c
}

func (c *Calculator) Subtract(a float64) *Calculator {
	c.result -= a
	return c
}

func (c *Calculator) Multiply(a float64) *Calculator {
	c.result *= a
	return c
}

func (c *Calculator) Divide(a float64) *Calculator {
	if a == 0 {
		panic("Деление на ноль!")
	}
	c.result /= a
	return c
}

func (c *Calculator) Power(n int) *Calculator {
	pow := math.Pow(float64(c.result), float64(n))
	c.result = pow
	return c
}

func main() {
	c := &Calculator{0}
	fmt.Println(*c.Add(5).Subtract(2).Multiply(3).Divide(4).Power(2))

}
