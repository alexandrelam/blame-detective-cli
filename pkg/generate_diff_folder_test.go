package pkg

import "testing"

func TestHasCommitHash(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"commit123", true},
		{"not a commit", false},
		{"commitxyz", true},
	}

	for _, tc := range testCases {
		result := hasCommitHash(tc.input)
		if result != tc.expected {
			t.Errorf("hasCommitHash(%s) expected %t, got %t", tc.input, tc.expected, result)
		}
	}
}

func TestHasFilePath(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"diff main.go", true},
		{" diff", false},
		{"sdlfkj different", false},
		{"diff_xyz", true},
	}

	for _, tc := range testCases {
		result := hasFilePath(tc.input)
		if result != tc.expected {
			t.Errorf("hasFilePath(%s) expected %t, got %t", tc.input, tc.expected, result)
		}
	}
}

func TestGetFilePath(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"diff --git a/README.md b/README.md", "README.md"},
		{"diff --git a/src/file.go b/src/file.go", "src/file.go"},
		{"diff --git a/foo/bar/baz.txt b/foo/bar/baz.txt", "foo/bar/baz.txt"},
		{"diff --git some/invalid/line", ""},
	}

	for _, tc := range testCases {
		result := getCommitFilePath(tc.input)
		if result != tc.expected {
			t.Errorf("getFilePath(%s) expected %s, got %s", tc.input, tc.expected, result)
		}
	}
}
