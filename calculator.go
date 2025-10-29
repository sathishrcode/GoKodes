package main

import "fmt"

func main() {
	var a, b float64
	var s string
	fmt.Println("enter two set of numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println("Choose 1-4 for desired operation: \n1. +  \n2. -  \n3. *  \n4. /  ")
	fmt.Scan(&s)
	switch s {
	//1234 == +-*/
	case "1":
		fmt.Println(a + b)
	case "2":
		fmt.Println(a - b)
	case "3":
		fmt.Println(a * b)
	case "4":
		fmt.Println(a / b)
	default:
		fmt.Println("err")
	}

}
