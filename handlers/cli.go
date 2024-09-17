package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Calculator interface{ Calculate(a, b int) int }

type CLIHandler struct {
	calculator Calculator
	output     io.Writer
}

func NewCLIHandler(calculator Calculator, output io.Writer) *CLIHandler {
	return &CLIHandler{calculator: calculator, output: output}
}

func (this *CLIHandler) Handle(args []string) error {
	if this.calculator == nil {
		return errNilCalculator
	}
	if len(args) != 2 {
		return fmt.Errorf("%w: two args required (you provided %d)", errTooFewArgs, len(args))
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: first arg (%s) %w", errMalformedArgument, args[0], err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: second arg (%s) %w", errMalformedArgument, args[1], err)
	}
	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintln(this.output, c)
	if err != nil {
		return fmt.Errorf("%w: %w", errOutputWriteErr, err)
	}
	return nil
}

var (
	errTooFewArgs        = errors.New("usage: calc <a> <b>")
	errMalformedArgument = errors.New("invalid argument")
	errOutputWriteErr    = errors.New("output writer err")
	errNilCalculator     = errors.New("calculator required")
)
