package utils

import "regexp"

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

// IsMD5 checks if the length of the string matches MD5 hash length.
func IsMD5(hash string) bool {
    return len(hash) == md5Length && IsHex(hash)
}

// IsSHA1 checks if the length of the string matches SHA1 hash length.
func IsSHA1(hash string) bool {
    return len(hash) == sha1Length && IsHex(hash)
}

// IsSHA256 checks if the length of the string matches SHA256 hash length.
func IsSHA256(hash string) bool {
    return len(hash) == sha256Length && IsHex(hash)
}

// IsSHA512 checks if the length of the string matches SHA512 hash length.
func IsSHA512(hash string) bool {
    return len(hash) == sha512Length && IsHex(hash)
}
