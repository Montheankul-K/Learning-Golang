package main

import "fmt"

func main() {
	// Variable : var name type = value
	var a int = 2
	var msg string = "Message"
	fmt.Println(a)
	fmt.Println(msg)

	// Dinamics type variable : name := value
	b := 2
	fmt.Println(b)

	// Convert type
	c := float64(a)
	fmt.Println(c)

	// Interface{} = any : มีค่าเป้นอะไรก็ได้
	var d any = "Second Message"
	fmt.Println(d)

	// Convert any to other type
	e, ok := d.(int)
	// ok : เช็คว่า convert ได้หรือไม่
	if !ok {
		panic(("Error"))
	}
	fmt.Println(e)
}
