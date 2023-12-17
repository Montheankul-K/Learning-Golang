package main

import "fmt"

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	// Pointer : เป็น variable รูปแบบหนึ่งที่เก็บทั้ง value และ address
	a := new(int) // ประกาศ pointer
	b := 10

	// Reference pointer a to b
	*a = b // หริือ a = &b

	// Print pointer ใช้ %p
	fmt.Printf("%p\n", a)  // Reference address
	fmt.Printf("%p\n", &b) // b address
	fmt.Printf("%p\n", &a) // a (pointer) address
	fmt.Printf("%d\n", *a) // Reference value

	x := 5
	y := 10
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
