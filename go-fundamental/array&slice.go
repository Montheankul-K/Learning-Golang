package main

import "fmt"

func arrayWatcher(data []int) {
	// Array ประกาศมาเป็น pointer จึงไม่ต้องใช้ return
	// Copy pointer แต่ละตัว value มาแต่ไม่ได้ copy pointer ของ slice มาจึงเปลี่ยนได้แค่ value ใน slice
	data[0] = 0
}

// []*int : slice of pointer
// *[]int : pointer of slice เก็บ address ของ slice

func arrayAppend(data *[]int) {
	*data = append(*data, 5)
}

func main() {
	// Array : name := [size]type{value}
	// Array ไม่ค่อยได้ใช้่ส่วนใหญ่จะใช้ slice
	a := [3]int{1, 2, 3}
	fmt.Println((a[1]))

	// Slice : เป็น array ที่ dynamic ไม่ fixed size
	b := []int{1, 2, 3}
	fmt.Println(b)

	// Append value : slice = append(slice,value)
	b = append(b, 4)
	fmt.Println(b)

	arrayWatcher(b)
	fmt.Println(b)

	arrayAppend(&b)
	fmt.Println(b)

	// Initial slice : name := make([]int, len, capacity) ปกติไม่ได้ใช้ capacity
	c := make([]int, 0)
	c = append(c, 4)
	fmt.Println(c)

	// Loop in array
	for i := range b {
		fmt.Println(a[i])
	}

	for _, value := range b {
		fmt.Println(value)
	}
}
