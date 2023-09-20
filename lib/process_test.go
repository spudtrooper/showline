package lib

import (
	"testing"
)

func TestGetDisplaySpecValid(t *testing.T) {
	tests := []struct {
		fileSpec    string
		expected    *displaySpec
		expectedErr error
	}{
		{
			fileSpec: "example.go:42",
			expected: &displaySpec{
				file: "example.go",
				line: 42,
			},
			expectedErr: nil,
		},
		{
			fileSpec: "another_file.txt:10",
			expected: &displaySpec{
				file: "another_file.txt",
				line: 10,
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.fileSpec, func(t *testing.T) {
			result, err := getDisplaySpec(test.fileSpec)

			if err != test.expectedErr {
				t.Errorf("Expected error: %v, got: %v", test.expectedErr, err)
			}

			if err == nil {
				if result.file != test.expected.file || result.line != test.expected.line {
					t.Errorf("Expected: %v, got: %v", test.expected, result)
				}
			}
		})
	}
}

func TestGetDisplaySpecInvalid(t *testing.T) {
	invalidFileSpecs := []string{
		"invalid_file",
		"file.go:",
		":42",
		"file.go:0",
		"file.go:-1",
		"file.go:abc",
	}

	for _, fileSpec := range invalidFileSpecs {
		t.Run(fileSpec, func(t *testing.T) {
			result, err := getDisplaySpec(fileSpec)

			if err == nil {
				t.Errorf("Expected error, got: %v", result)
			}
		})
	}
}
