package main

import (
	"flag"
	"os"

	"github.com/mdwhatcott/calc-apps/handlers"
	"github.com/mdwhatcott/calc-lib/calc"
)

func main() {
	var op string
	flag.StringVar(&op, "op", "+", "One of + - * /")
	flag.Parse()
	handler := handlers.NewCLIHandler(calculators[op], os.Stdout)
	err := handler.Handle(flag.Args())
	if err != nil {
		panic(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": calc.Addition{},
	"-": calc.Subtraction{},
	"*": calc.Multiplication{},
	"/": calc.Division{},
}
