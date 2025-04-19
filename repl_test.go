package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  CHarmander BUlbasaur PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "MITCH NOAH",
			expected: []string{"mitch", "noah"},
		},
		{
			input: "   MITCH NOAH   ",
			expected: []string{"mitch", "noah"},
		},
		{
			input: "mitch noah",
			expected: []string{"mitch", "noah"},
		},
		{
			input: "   MiTcH nOaH     ",
			expected: []string{"mitch", "noah"},
		},
		{
			input: "      ",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual: %v does not match length of expected: %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word: %v does not equal expected word: %v", word, expectedWord)
			}
		}
	}
}