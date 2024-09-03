package handlers

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/mdwhatcott/calc-lib/calc"
)

func assertEqual(t *testing.T, expected, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Helper()
		t.Errorf("\n"+
			"expected: %v\n"+
			"actual:   %v", expected, actual)
	}
}
func assertError(t *testing.T, expected, actual error) {
	if !errors.Is(actual, expected) {
		t.Helper()
		t.Errorf("expected [%v], got [%v]", expected, actual)
	}
}

func TestTooFewArguments(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewCLIHandler(calc.Addition{}, &output)
	err := handler.Handle(nil)
	assertError(t, errTooFewArgs, err)
	assertEqual(t, "", output.String())
}
func TestInvalidFirstArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewCLIHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"NaN", "1"})
	assertError(t, errMalformedArgument, err)
	assertEqual(t, "", output.String())
}
func TestInvalidSecondArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewCLIHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "NaN"})
	assertError(t, errMalformedArgument, err)
	assertEqual(t, "", output.String())
}
func TestOutputWriterError(t *testing.T) {
	taco := errors.New("taco")
	output := &ErringWriter{err: taco}
	handler := NewCLIHandler(calc.Addition{}, output)
	err := handler.Handle([]string{"1", "2"})
	assertError(t, errOutputWriteErr, err)
	assertError(t, taco, err)
}
func TestHappyPath(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewCLIHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "2"})
	assertEqual(t, nil, err)
	assertEqual(t, "3\n", output.String())
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write([]byte) (n int, err error) {
	return 0, this.err
}
