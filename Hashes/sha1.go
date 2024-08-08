package hashes

// IsSHA1 checks if the hash is a SHA-1 hash.
func IsSHA1(hash string) bool {
    return len(hash) == 40
}
