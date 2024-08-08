package hashes

import (
	"regexp"
)

// Define the length of a SHA-256 hash in hexadecimal form
const sha256Length = 64

// IsHex checks if a string is a valid hexadecimal string
func IsHex(s string) bool {
	matched, _ := regexp.MatchString("^[a-fA-F0-9]+$", s)
	return matched
}

// IsSHA256 checks if the hash is a valid SHA-256 hash
func IsSHA256(hash string) bool {
	return len(hash) == sha256Length && IsHex(hash)
}
