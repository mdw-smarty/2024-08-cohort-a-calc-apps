package gunit

import (
	"reflect"
	"strings"
	"testing"
)

func Run(t *testing.T, fixture any) {
	fixtureValue := reflect.ValueOf(fixture)
	fixtureType := fixtureValue.Type().Elem()
	var testNames []string
	var hasSetup bool
	for i := 0; i < fixtureValue.Type().NumMethod(); i++ {
		method := fixtureValue.Type().Method(i)
		if strings.HasPrefix(method.Name, "Test") {
			testNames = append(testNames, method.Name)
		}
		if strings.HasPrefix(method.Name, "Setup") {
			hasSetup = true
		}
	}
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			instance := reflect.New(fixtureType)
			instance.Elem().FieldByName("Fixture").Set(reflect.ValueOf(&Fixture{T: t}))
			if hasSetup {
				instance.MethodByName("Setup").Call(nil)
			}
			instance.MethodByName(name).Call(nil)
		})
	}
}

type Fixture struct{ *testing.T }

func (this *Fixture) So(actual any, assert assertion, expected ...any) bool {
	err := assert(actual, expected...)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
	return err == nil
}

type assertion func(actual any, expected ...any) error
