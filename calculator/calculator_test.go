package calculator_test

import (
	"testing"

	"github.com/practice/calculator"
)

type TestCase struct {
	value    int
	expected int
	actual   int
}

func TestCalculatePows(t *testing.T) {
	testcase := TestCase{
		expected: 16,
		value:    4,
	}
	testcase.actual = calculator.CalculatePows(testcase.value)
	if testcase.actual != testcase.expected {
		t.Fail()
	}
}

func TestNegativeCalculatePows(t *testing.T) {
	testcase := TestCase{
		expected: 16,
		value:    -4,
	}
	testcase.actual = calculator.CalculatePows(testcase.value)
	if testcase.actual != testcase.expected {
		t.Fail()
	}
}
