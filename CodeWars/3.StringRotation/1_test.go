package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Log("hey boy")
}

type TestCase struct {
	input    int
	expected int
}

func TestSqrt(t *testing.T) {
	testCases := []TestCase{
		{input: 18, expected: 2},
		{input: 101, expected: 2},
		{input: 10, expected: 2},
		{input: 2, expected: 1},
		{input: 1, expected: 1},
		{input: 4, expected: 1},
	}
	for _, test := range testCases {
		got := Clean(test.input)
		if got != test.expected {
			t.Errorf("input: %d, expected: %d, got:%d \n", test.input, test.expected, got)
		}
	}
}
