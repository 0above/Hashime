package hashes

// Define the lengths of SHA-3 hashes in hexadecimal form
const (
    sha3_224Length = 56
    sha3_256Length = 64
    sha3_384Length = 96
    sha3_512Length = 128
)

// IsSHA3_224 checks if the hash is a valid SHA3-224 hash
func IsSHA3_224(hash string) bool {
    return len(hash) == sha3_224Length && IsHex(hash)
}

// IsSHA3_256 checks if the hash is a valid SHA3-256 hash
func IsSHA3_256(hash string) bool {
    return len(hash) == sha3_256Length && IsHex(hash)
}

// IsSHA3_384 checks if the hash is a valid SHA3-384 hash
func IsSHA3_384(hash string) bool {
    return len(hash) == sha3_384Length && IsHex(hash)
}

// IsSHA3_512 checks if the hash is a valid SHA3-512 hash
func IsSHA3_512(hash string) bool {
    return len(hash) == sha3_512Length && IsHex(hash)
}
