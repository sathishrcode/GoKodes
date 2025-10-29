package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		fmt.Print(rand.Intn(6)+1, " ") //+1 coz we dont need 0
	}
}
