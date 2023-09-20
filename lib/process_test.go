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
			res, err := getDisplaySpec(test.fileSpec)

			if err != test.expectedErr {
				t.Errorf("Expected error: %v, got: %v", test.expectedErr, err)
			}

			if err == nil {
				if res.File != test.expected.File || res.Line != test.expected.Line {
					t.Errorf("Expected: %v, got: %v", test.expected, res)
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
			if res, err := getDisplaySpec(fileSpec); err == nil {
				t.Errorf("Expected error, got: %v for fileSpec: %s", res, fileSpec)
			}
		})
	}
}
