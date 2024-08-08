package hashes

// Define the lengths of FNV hashes in hexadecimal form
const (
    fnv1Length  = 8
    fnv1aLength = 8
)

// IsFNV1 checks if the hash is a valid FNV-1 hash
func IsFNV1(hash string) bool {
    return len(hash) == fnv1Length && IsHex(hash)
}

// IsFNV1a checks if the hash is a valid FNV-1a hash
func IsFNV1a(hash string) bool {
    return len(hash) == fnv1aLength && IsHex(hash)
}
