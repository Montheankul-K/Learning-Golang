package main

import "fmt"

type Airplane struct {
	Name string
}

func (a Airplane) Fly() {
	fmt.Printf("%s: Flying\n", a.Name)
}

func main() {
	airbus := Airplane{Name: "Airbus"}
	boeing := Airplane{Name: "Boeing"}

	airbus.Fly()
	boeing.Fly()
}
