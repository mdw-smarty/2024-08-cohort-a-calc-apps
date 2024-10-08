package should

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

var AssertionFailure = errors.New("assertion failure")

func BeBlank(actual any, expected ...any) error {
	if len(expected) != 0 {
		return fmt.Errorf("expected no expected values, got %d", len(expected))
	}
	return Equal(actual, "")
}
func BeNil(actual any, expected ...any) error {
	if len(expected) != 0 {
		return fmt.Errorf("expected no expected values, got %d", len(expected))
	}
	if actual == nil {
		return nil
	}
	return fmt.Errorf("%w: expected nil, got: %v", AssertionFailure, actual)
}

func BeError(actual any, expected ...any) error {
	if len(expected) != 1 {
		return fmt.Errorf("expected a single value, got %d", len(expected))
	}
	actualErr, ok := actual.(error)
	if !ok {
		return fmt.Errorf("expected an error, got %T", actual)
	}
	expectedErr, ok := expected[0].(error)
	if !ok {
		return fmt.Errorf("expected an error, got %T", expected[0])
	}
	if !errors.Is(actualErr, expectedErr) {
		return fmt.Errorf(
			"%w: expected %v to wrap %v, but it didn't",
			AssertionFailure, actual, expected[0],
		)
	}
	return nil
}

func Equal(actual any, expected ...any) error {
	if len(expected) != 1 {
		return fmt.Errorf("expected a single value, got %d", len(expected))
	}
	if !reflect.DeepEqual(expected[0], actual) {
		return fmt.Errorf("%w\n"+
			"expected: %v\n"+
			"actual:   %v",
			AssertionFailure,
			expected,
			actual,
		)
	}
	return nil
}

type Assertion func(actual any, expected ...any) error

func So(t *testing.T, actual any, assertion Assertion, expected ...any) bool {
	err := assertion(actual, expected...)
	if err != nil {
		t.Helper()
		t.Error(err)
	}
	return err == nil
}
