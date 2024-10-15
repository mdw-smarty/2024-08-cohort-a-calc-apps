package gunit_test

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/calc-apps/externals/gunit"
	"github.com/mdwhatcott/calc-apps/externals/should"
)

func Test(t *testing.T) {
	gunit.Run(t, new(TestingFixture))
}

type TestingFixture struct {
	*gunit.Fixture
}

func (this *TestingFixture) Setup() {
	fmt.Println("setup")
}
func (this *TestingFixture) TestSomething() {
	fmt.Println("Hello Cohort A")
	this.So(1, should.Equal, 1)
}
func (this *TestingFixture) TestSomethingElse() {
	fmt.Println("Hello, Ryan", this)
}
