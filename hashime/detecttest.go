package hashime

import (
	"testing"
)

func TestDetectHashType(t *testing.T) {
	tests := []struct {
		hash     string
		expected HashType
	}{
		{"d41d8cd98f00b204e9800998ecf8427e", MD5},      // MD5
		{"a9993e364706816aba3e25717850c26d", SHA1},    // SHA1
		{"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", SHA256}, // SHA256
		{"cf83e1357eefb8bd9d61e0d6e5ba4d2793a8a9a5cfc5418f1d254b3e8e4e0e3b29e8b5b7fdd542a3b8a6c44c8f4c34a4", SHA512}, // SHA512
		{"invalidhash", Unknown},
	}

	for _, test := range tests {
		result := DetectHashType(test.hash)
		if result != test.expected {
			t.Errorf("DetectHashType(%q) = %v; want %v", test.hash, result, test.expected)
		}
	}
}
