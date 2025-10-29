package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(100)
	var g int
	fmt.Println("i have a number in MIND, you guess until its right")
	for {
		fmt.Scan(&g)
		if g == n {
			fmt.Println("correct one")
			break
		}
		if g < n {
			fmt.Println("UP")
		} else {
			fmt.Println("DOWN")
		}
	}
}
