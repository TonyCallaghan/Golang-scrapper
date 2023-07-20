package core

import (
	"testing"
)

func TestCleanupDescription(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"Lots   of spaces  in this     sentence.", "Lots of spaces in this sentence."},
		{"Test1.Test2?Test3,Test4!Done", "Test1. Test2? Test3, Test4! Done"},
		{"Test1.    Test2", "Test1. Test2"},
	}

	for _, test := range tests {
		if output := CleanupDescription(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
