package hashime

import "github.com/0above/Hashime/hashes"

// HashType represents different types of hashes.
type HashType int

const (
    Unknown HashType = iota
    MD5
    SHA1
    SHA256
    SHA512
    // Add other hash types as needed
)

// DetectHashType detects the type of hash based on the hash value.
func DetectHashType(hash string) HashType {
    if hashes.IsHex(hash) {
        switch {
        case hashes.IsMD5(hash):
            return MD5
        case hashes.IsSHA1(hash):
            return SHA1
        case hashes.IsSHA256(hash):
            return SHA256
        case hashes.IsSHA512(hash):
            return SHA512
        }
    }
    return Unknown
}
