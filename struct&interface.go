package main

import (
	"encoding/json"
	"fmt"
)

// Struct : เป็น data type ประเภทหนึ่ง
type Book struct {
	// ถ้าตั้งชื่อเป็น lower case จะเป็น private struct
	// ถ้า upper case จะเป็น public struct import โดย package.struct{}
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	// `json:"field"` เป็นการกำหนดชื่อ field ให้ json
}

func Debug(obj interface{}) {
	raw, _ := json.MarshalIndent(&obj, "", "\t") // Convert to json
	// ใส่ _ ถ้่าไม่อยากรับค่า error
	// เวลาใช้ MarshalIndent() ให้ pass เป็น address
	fmt.Println(string(raw))
}

type book struct {
	Book
}

func (b *book) GetBook() *book {
	return b
}

// เวลาใช้งานจริง function นิยมรับ param เป็น pointer
func (b *book) SetTitle(title string) {
	b.Title = title
}

// Interface : มีประโยชน์เวลาเขียนแบบ OOP
type IBook interface { // นิยมขึ้นต้นด้วย I เวลาตั้งชื่อ interface
	GetBook() *book
	SetTitle(title string)
}

func NewBook(title, author string) IBook {
	return &book{
		Book{
			ID:     1,
			Title:  title,
			Author: author,
		}, // เอา book ไปเข้าถึง interface
	}
}

func main() {
	a := Book{
		ID:     1,
		Title:  "Golang",
		Author: "Gopher",
	}

	b := book{a}
	Debug(b.GetBook())

	b.SetTitle("Golang Advance")
	Debug(b.GetBook())

	// เหมือนการสร้าง Object
	c := NewBook("Java Basic", "-")
	Debug(c.GetBook())
}
