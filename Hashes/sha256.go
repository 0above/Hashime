package hashes

// IsSHA256 checks if the hash is a SHA-256 hash.
func IsSHA256(hash string) bool {
    return len(hash) == 64
}
