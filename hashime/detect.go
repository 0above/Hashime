package hashime

import "github.com/0above/Hashime/utils"

// HashType represents different types of hashes.
type HashType int

const (
    Unknown HashType = iota
    MD5
    SHA1
    SHA256
    SHA512
)

// DetectHashType detects the type of hash based on the hash value.
func DetectHashType(hash string) HashType {
    if utils.IsHex(hash) {
        switch {
        case utils.IsMD5(hash):
            return MD5
        case utils.IsSHA1(hash):
            return SHA1
        case utils.IsSHA256(hash):
            return SHA256
        case utils.IsSHA512(hash):
            return SHA512
        }
    }
    return Unknown
}
