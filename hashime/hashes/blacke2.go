package hashes

// Define the length of a Blake2b hash in hexadecimal form
const blake2bLength = 64

// IsBlake2b checks if the hash is a valid Blake2b hash
func IsBlake2b(hash string) bool {
    return len(hash) == blake2bLength && IsHex(hash)
}
