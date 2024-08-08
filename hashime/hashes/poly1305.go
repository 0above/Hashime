package hashes

// Define the length of a Poly1305 hash in hexadecimal form
const poly1305Length = 32

// IsPoly1305 checks if the hash is a valid Poly1305 hash
func IsPoly1305(hash string) bool {
    return len(hash) == poly1305Length && IsHex(hash)
}
