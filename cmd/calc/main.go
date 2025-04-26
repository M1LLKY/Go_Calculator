package main

import (
	"Calculator/internal/mathlib"
	"fmt"
)

func main() {
	fmt.Println("Please, enter number of operation you want to execute:\n1. Addition (+)\n2. Substraction (-)\n3. Multiplication (*)\n4. Division (/)")
	var operation int
	fmt.Scan(&operation)
	fmt.Print("Now enter two numbers to work with: ")
	var num_1, num_2 int
	fmt.Scan(&num_1, &num_2)
	switch operation {
	case 1:
		fmt.Println(num_1, "+", num_2, "=", mathlib.Add(num_1, num_2))
	case 2:
		fmt.Println(num_1, "-", num_2, "=", mathlib.Sub(num_1, num_2))
	case 3:
		fmt.Println(num_1, "*", num_2, "=", mathlib.Multi(num_1, num_2))
	case 4:
		fmt.Println(num_1, "/", num_2, "=", mathlib.Div(num_1, num_2))
	}
}
