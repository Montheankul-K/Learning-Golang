package main

import "fmt"

type (
	Keyboard interface {
		Input()
	}

	MechanicalKeyboard struct {
		SwitchType string
		Size       string
		OS         string
	}

	NormalKeyboard struct {
		IsCanPress bool
	}
)

func (m MechanicalKeyboard) Input() {
	fmt.Println("Pressing the key on the mechanical keyboard: ", m.SwitchType, m.Size, m.OS)
}

func (n NormalKeyboard) Input() {
	fmt.Println("Pressing the key on the normal keyboard")
}

func main() {
	mechanicalKeyboard := MechanicalKeyboard{
		SwitchType: "Blue switch",
		Size:       "60%",
		OS:         "Windows, Mac OS",
	}

	normalKeyboard := NormalKeyboard{
		IsCanPress: true,
	}

	mechanicalKeyboard.Input()
	normalKeyboard.Input()
}
