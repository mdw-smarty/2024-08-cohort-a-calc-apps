package handlers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/mdwhatcott/calc-apps/externals/gunit"
	"github.com/mdwhatcott/calc-apps/externals/should"
	"github.com/mdwhatcott/calc-lib/calc"
)

func TestCLIFixture(t *testing.T) {
	gunit.Run(t, new(CLIFixture))
}

type CLIFixture struct {
	*gunit.Fixture
	output  bytes.Buffer
	handler *CLIHandler
}

func (this *CLIFixture) Setup() {
	this.handler = NewCLIHandler(calc.Addition{}, &this.output)
}

func (this *CLIFixture) TestTooFewArguments() {
	err := this.handler.Handle(nil)
	this.So(err, should.BeError, errTooFewArgs)
	this.So(this.output.String(), should.BeBlank)
}
func (this *CLIFixture) TestNilCalculator() {
	handler := NewCLIHandler(nil, nil)
	err := handler.Handle(nil)
	this.So(err, should.BeError, errNilCalculator)
}
func (this *CLIFixture) TestInvalidFirstArg() {
	err := this.handler.Handle([]string{"NaN", "1"})
	this.So(err, should.BeError, errMalformedArgument)
	this.So(this.output.String(), should.BeBlank)
}
func (this *CLIFixture) TestInvalidSecondArg() {
	err := this.handler.Handle([]string{"1", "NaN"})
	this.So(err, should.BeError, errMalformedArgument)
	this.So(this.output.String(), should.BeBlank)
}
func (this *CLIFixture) TestOutputWriterError() {
	taco := errors.New("taco")
	output := &ErringWriter{err: taco}
	this.handler = NewCLIHandler(calc.Addition{}, output)
	err := this.handler.Handle([]string{"1", "2"})
	this.So(err, should.BeError, errOutputWriteErr)
	this.So(err, should.BeError, taco)
}
func (this *CLIFixture) TestHappyPath() {
	err := this.handler.Handle([]string{"1", "2"})
	this.So(err, should.BeNil)
	this.So(this.output.String(), should.Equal, "3\n")
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write([]byte) (n int, err error) {
	return 0, this.err
}
