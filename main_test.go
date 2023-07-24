// main_test.go
package main

import (
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input   string
		words   bool
		expWord int
		expLine int
	}{
		{
			input:   "This is a sample text\n with multiple lines\n exit",
			words:   true,
			expWord: 8,
			expLine: 3,
		},
		{
			input:   "Another example with fewer words",
			words:   true,
			expWord: 5,
			expLine: 1,
		},
		{
			input:   "Single line\n exit",
			words:   true,
			expWord: 2,
			expLine: 2,
		},
		{
			input:   "A\nB\nC\n exit",
			words:   false,
			expWord: 4,
			expLine: 3,
		},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.input)
		wordCount := count(reader, test.words)
		if wordCount != test.expWord {
			t.Errorf("Word count mismatch for input '%s'. Expected: %d, Got: %d", test.input, test.expWord, wordCount)
		}

		reader = strings.NewReader(test.input)
		lineCount := count(reader, !test.words)
		if lineCount != test.expLine {
			t.Errorf("Line count mismatch for input '%s'. Expected: %d, Got: %d", test.input, test.expLine, lineCount)
		}
	}
}

func TestStatement(t *testing.T) {
	ch := make(chan string)
	go statement(true, ch)
	msg := <-ch
	if msg != "The words count is" {
		t.Errorf("Statement mismatch. Expected: 'The words count is', Got: '%s'", msg)
	}

	go statement(false, ch)
	msg = <-ch
	if msg != "The lines count is" {
		t.Errorf("Statement mismatch. Expected: 'The lines count is', Got: '%s'", msg)
	}
}
