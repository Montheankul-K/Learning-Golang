package main

import "fmt"

func main() {
	a := 2
	if a != 2 {
		fmt.Println("Is not 2")
	} else if a == 2 {
		fmt.Println(a)
	} else {
		fmt.Println("Message")
	}
}
