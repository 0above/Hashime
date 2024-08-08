package hashes

// Define the length of a SHA-1 hash in hexadecimal form
const sha1Length = 40

// IsSHA1 checks if the hash is a valid SHA-1 hash
func IsSHA1(hash string) bool {
    return len(hash) == sha1Length && IsHex(hash)
}
