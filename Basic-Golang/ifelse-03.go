package main

import (
	"errors"
)

func genError() error {
	return errors.New("Some error")
}

func main() {
	// Defind variable in if condition : variable := value; condition
	// Vriable จบใน condition เท่านั้น
	if err := genError(); err != nil {
		panic(err.Error())
	}
}
