package hashes

import (
	"regexp"
)

// Constants for hash lengths
const (
	md5Length    = 32
	sha1Length   = 40
	sha256Length = 64
	sha512Length = 128
)

// IsHex checks if a string is a valid hexadecimal string.
func IsHex(s string) bool {
	matched, _ := regexp.MatchString("^[a-fA-F0-9]+$", s)
	return matched
}

// IsMD5 checks if the hash is an MD5 hash.
func IsMD5(hash string) bool {
	return len(hash) == md5Length && IsHex(hash)
}

// IsSHA1 checks if the hash is a SHA1 hash.
func IsSHA1(hash string) bool {
	return len(hash) == sha1Length && IsHex(hash)
}

// IsSHA256 checks if the hash is a SHA256 hash.
func IsSHA256(hash string) bool {
	return len(hash) == sha256Length && IsHex(hash)
}

// IsSHA512 checks if the hash is a SHA512 hash.
func IsSHA512(hash string) bool {
	return len(hash) == sha512Length && IsHex(hash)
}
