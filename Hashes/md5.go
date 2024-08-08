package hashes

// IsMD5 checks if the hash is an MD5 hash.
func IsMD5(hash string) bool {
    return len(hash) == 32
}
