package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	t := []string{} //tasks
	s := bufio.NewScanner(os.Stdin)
	//for menu/CRUD
	for {
		fmt.Println("1.add\n2 view \n3. Delete")
		s.Scan()
		switch s.Text() {
		case "1":
			s.Scan()
			t = append(t, s.Text())
		case "2":
			fmt.Println("tasks are below")
			for i, v := range t {
				fmt.Println(i+1, v)
			}
		case "3":
			return
		}
	}
}
