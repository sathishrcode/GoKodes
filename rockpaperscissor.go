package main

import (
	"fmt"
	"math/rand"
)

func main() {

	r := []string{"r", "p", "s"}
	var a string //user ip
	fmt.Println("choose one among r/p/s  ")
	fmt.Scan(&a)
	b := r[rand.Intn(3)]
	fmt.Println(b)
	if a == b {
		fmt.Println("its a DRAW")
	} else if a == "r" && b == "s" || a == "p" && b == "r" || a == "s" && b == "p" {
		fmt.Println("you win")
	} else {
		fmt.Println("u loose")
	}
}
