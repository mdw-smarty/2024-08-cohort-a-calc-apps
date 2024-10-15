package handlers

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/mdwhatcott/calc-lib/calc"
	"github.com/smarty/assertions/should"
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
	should.So(t, err, should.Wrap, errTooFewArgs)
	should.So(t, output.String(), should.BeBlank)
}
func TestNilCalculator(t *testing.T) {
	handler := NewCLIHandler(nil, nil)
	err := handler.Handle(nil)
	should.So(t, err, should.Wrap, errNilCalculator)
}
func TestInvalidFirstArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewCLIHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"NaN", "1"})
	should.So(t, err, should.Wrap, errMalformedArgument)
	should.So(t, output.String(), should.BeBlank)
}
func TestInvalidSecondArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewCLIHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "NaN"})
	should.So(t, err, should.Wrap, errMalformedArgument)
	should.So(t, output.String(), should.BeBlank)
}
func TestOutputWriterError(t *testing.T) {
	taco := errors.New("taco")
	output := &ErringWriter{err: taco}
	handler := NewCLIHandler(calc.Addition{}, output)
	err := handler.Handle([]string{"1", "2"})
	should.So(t, err, should.Wrap, errOutputWriteErr)
	should.So(t, err, should.Wrap, taco)
}
func TestHappyPath(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewCLIHandler(calc.Addition{}, &output)

	err := handler.Handle([]string{"1", "2"})

	should.So(t, err, should.BeNil)
	should.So(t, output.String(), should.Equal, "3\n")
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write([]byte) (n int, err error) {
	return 0, this.err
}
