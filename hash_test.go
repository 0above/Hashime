// File: detecttest.go
package hashime_test

import (
	"testing"

	"github.com/0above/Hashime/hashime" // Adjust import path if needed
)

// Define test cases
var testCases = []struct {
	hash     string
	expected hashime.HashType
}{
	// Test cases for various hash types
	{"5f4dcc3b5aa765d61d8327deb882cf99", hashime.MD5}, // MD5 hash of "password"
	{"5baa61e4c9b93f3f0682250b6cf8331b4eea4f8b", hashime.SHA1}, // SHA1 hash of "password"
	{"5e884898da28047151d0e56f8dc6292773603d0d1b4d4d80f6b5f167e9c5b6f8", hashime.SHA256}, // SHA256 hash of "password"
	{"cf83e1357eefb8bdc3e1d8f1e7f8f5b9e2b7ffb3b7b7f7d5d44f618b2f835f8a9efc6d6f4b0d0a7e7b1a00b24c0fbe6bdeab2754b4c1f945f0e95d3d065fc82d", hashime.SHA512},
	{"00000000", hashime.Unknown}, // Invalid hash
}

func TestDetectHashType(t *testing.T) {
	for _, tc := range testCases {
		result := hashime.DetectHashType(tc.hash)
		if result != tc.expected {
			t.Errorf("DetectHashType(%s) = %v; want %v", tc.hash, result, tc.expected)
		}
	}
}
