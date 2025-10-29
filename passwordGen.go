package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := "qwertyuiopasdfghjjklzxcvbnm"
	b := "QWERTYUIOPLKJHGFDSAZXCVBNM"
	c := "1234567890!@#$%^&*"
	s := a+b+c
	for i:=0; i<10; i++ {
		fmt.Printf("%c", s[rand.Intn(len(s))])
	}
}
