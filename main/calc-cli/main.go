package main

import (
	"os"

	"github.com/mdwhatcott/calc-apps/handlers"
	"github.com/mdwhatcott/calc-lib/calc"
)

func main() {
	handler := handlers.NewCLIHandler(calc.Addition{}, os.Stdout)
	err := handler.Handle(os.Args[1:])
	if err != nil {
		panic(err)
	}
}
