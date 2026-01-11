package strutil

import (
	"reflect"
	"testing"
)

func TestCaseSetMaker(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected CaseSet
	}{
		{
			name:  "PascalCase input",
			input: "HelloWorld",
			expected: CaseSet{
				Kebab:          "hello-world",
				Snake:          "hello_world",
				Pascal:         "HelloWorld",
				Camel:          "helloWorld",
				ScreamingSnake: "HELLO_WORLD",
			},
		},
		{
			name:  "snake_case input",
			input: "hello_world",
			expected: CaseSet{
				Kebab:          "hello-world",
				Snake:          "hello_world",
				Pascal:         "HelloWorld",
				Camel:          "helloWorld",
				ScreamingSnake: "HELLO_WORLD",
			},
		},
		{
			name:  "kebab-case input",
			input: "hello-world",
			expected: CaseSet{
				Kebab:          "hello-world",
				Snake:          "hello_world",
				Pascal:         "HelloWorld",
				Camel:          "helloWorld",
				ScreamingSnake: "HELLO_WORLD",
			},
		},
		{
			name:  "camelCase input",
			input: "helloWorld",
			expected: CaseSet{
				Kebab:          "hello-world",
				Snake:          "hello_world",
				Pascal:         "HelloWorld",
				Camel:          "helloWorld",
				ScreamingSnake: "HELLO_WORLD",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CaseSetMaker(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("CaseSetMaker() = %v, want %v", got, tt.expected)
			}
		})
	}
}
