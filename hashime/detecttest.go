package hashime

import (
	"testing"
)

func TestDetectHashType(t *testing.T) {
    tests := []struct {
        hash     string
        expected HashType
    }{
        {"d41d8cd98f00b204e9800998ecf8427e", MD5},         // MD5 hash
        {"a9993e364706816aba3e25717850c26d9c2f29a3", SHA1}, // SHA1 hash
        {"e99a18c428cb38d5f260853678922e03abdce8c", SHA256}, // SHA256 hash
        {"cf83e1357eefb8bd1c1b70e49b93e6efc0a1a8c8b41d2d1f2a5c2f6a9e5f4d7b4a0a4e4a3e8f6e5c084b0f94b4fbbf21", SHA512}, // SHA512 hash
        {"not_a_hash", Unknown},                            // Invalid hash
    }

    for _, tt := range tests {
        result := DetectHashType(tt.hash)
        if result != tt.expected {
            t.Errorf("DetectHashType(%s) = %v; want %v", tt.hash, result, tt.expected)
        }
    }
}
