package hashes

// IsSHA512 checks if the hash is a SHA-512 hash.
func IsSHA512(hash string) bool {
    return len(hash) == 128
}
