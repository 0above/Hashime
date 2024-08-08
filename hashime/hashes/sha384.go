package hashes

// Define the length of a SHA-384 hash in hexadecimal form
const sha384Length = 96

// IsSHA384 checks if the hash is a valid SHA-384 hash
func IsSHA384(hash string) bool {
    return len(hash) == sha384Length && IsHex(hash)
}
