package main

import "fmt"

func sum(msg string, a, b int) (string, int) {
	// ถ้า type ของ params เหมือนกันไม่ต้องประกาศ type ซ้ำ
	// function ที่เป็น global จะตั้งชื่อขึ้นต้นด้วย upper case
	return msg, a + b
}

func changer(x int) {
	// params x เป็น shadow variable หรือ local variable
	x = 20
}

func main() {
	fmt.Println(sum("Sum result is", 10, 20)) // Sum result is 30

	x := 10
	changer(10)
	fmt.Println(x)

	// เรียกใช้ global function : package.function()
}
