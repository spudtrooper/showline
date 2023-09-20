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
				File: "example.go",
				Line: 42,
			},
			expectedErr: nil,
		},
		{
			fileSpec: "another_file.txt:10",
			expected: &displaySpec{
				File: "another_file.txt",
				Line: 10,
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.fileSpec, func(t *testing.T) {
			var result displaySpec
			err := getDisplaySpec(test.fileSpec, &result)

			if err != test.expectedErr {
				t.Errorf("Expected error: %v, got: %v", test.expectedErr, err)
			}

			if err == nil {
				if result.File != test.expected.File || result.Line != test.expected.Line {
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
			var result displaySpec
			err := getDisplaySpec(fileSpec, &result)

			if err == nil {
				t.Errorf("Expected error, got: %v for fileSpec: %s", result, fileSpec)
			}
		})
	}
}
