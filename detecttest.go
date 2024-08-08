package hashime

import "testing"

func TestDetectHashType(t *testing.T) {
    tests := []struct {
        hash     string
        expected HashType
    }{
        {"5d41402abc4b2a76b9719d911017c592", MD5},
        {"a5c8c229f8b17f8d9465f9484b317e0985a1fa5b", SHA1},
        {"e99a18c428cb38d5f260853678922e03abd6e2f", SHA256},
        {"cf83e1357eefb8bd3a4d0d8a9d64e1d0a27a1c5f9a3f2d70c457a36d0a46d8b1b26860d6d4fd0422efb1a1e9f4d8d3d080c6e8d647e4", SHA512},
        {"invalidhash", Unknown},
    }

    for _, test := range tests {
        result := DetectHashType(test.hash)
        if result != test.expected {
            t.Errorf("DetectHashType(%s) = %v; want %v", test.hash, result, test.expected)
        }
    }
}
