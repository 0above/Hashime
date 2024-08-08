package hashes

// Define the length of a BLAKE3 hash in hexadecimal form
const blake3Length = 64

// IsBLAKE3 checks if the hash is a valid BLAKE3 hash
func IsBLAKE3(hash string) bool {
    return len(hash) == blake3Length && IsHex(hash)
}
