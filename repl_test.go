package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestCleanInput(t *testing.T) {
	fmt.Println("Testing")
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string {"hello", "world"},
		},
		{
			input: "this is for       middle spaces",
			expected: []string {"this", "is", "for", "middle", "spaces"},
		},
		{
			input: "      multiple leading",
			expected: []string {"multiple", "leading"},
		},
		{
			input: "multiple trailing        ",
			expected: []string{"multiple", "trailing"},
		},
		{
			input: "ThiS ISSSss for CapITalIZaTiON",
			expected: []string{"this", "isssss", "for", "capitalization"},
		},
		{
			input: "This    is   TO    MIX  EverYTHing     ",
			expected: []string{"this", "is", "to", "mix", "everything"},
		},
		
		// Add more cases
	}
	
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Fatalf("Expected len: %v, got: %v", len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if !reflect.DeepEqual(word, expectedWord) {
				t.Fatalf("Expected: %v, got: %v", word, expectedWord)
			}
		}
	}
}
