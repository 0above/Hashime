package hashes

// Define the length of a SHA-512 hash in hexadecimal form
const sha512Length = 128

// IsSHA512 checks if the hash is a valid SHA-512 hash
func IsSHA512(hash string) bool {
    return len(hash) == sha512Length && IsHex(hash)
}
