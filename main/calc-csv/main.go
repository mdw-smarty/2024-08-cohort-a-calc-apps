package main

import (
	"log"
	"os"

	"github.com/mdwhatcott/calc-apps/handlers"
	"github.com/mdwhatcott/calc-lib/calc"
)

func main() {
	handler := handlers.NewCSVHandler(os.Stdin, os.Stdout, calculators)
	err := handler.Handle()
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": calc.Addition{},
	"-": calc.Subtraction{},
	"*": calc.Multiplication{},
	"/": calc.Division{},
}
