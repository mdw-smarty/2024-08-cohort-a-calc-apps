package handlers

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/mdwhatcott/calc-lib/calc"
)

const input = `1,+,2
45,+
2,-,1
NaN,+,2
1,+,NaN
1,nop,2
3,*,4
20,/,10`

const expectedOutput = `1,+,2,3
2,-,1,1
3,*,4,12
20,/,10,2
`

func TestCSVHappyPath(t *testing.T) {
	var output bytes.Buffer

	err := NewCSVHandler(strings.NewReader(input), &output, calculators).Handle()

	assertError(t, nil, err)
	assertEqual(t, expectedOutput, output.String())
}

func TestCSVReadError(t *testing.T) {
	var input ErringReader
	boink := errors.New("boink")
	input.err = boink
	var output bytes.Buffer

	err := NewCSVHandler(input, &output, calculators).Handle()

	assertError(t, boink, err)
	assertEqual(t, "", output.String())
}

func TestCSVWriteError(t *testing.T) {
	var output ErringWriter
	boink := errors.New("boink")
	output.err = boink
	err := NewCSVHandler(strings.NewReader(input), &output, calculators).Handle()
	assertError(t, boink, err)
}

type ErringReader struct {
	err error
}

func (this ErringReader) Read([]byte) (n int, err error) {
	return 0, this.err
}

var calculators = map[string]Calculator{
	"+": calc.Addition{},
	"-": calc.Subtraction{},
	"*": calc.Multiplication{},
	"/": calc.Division{},
}
