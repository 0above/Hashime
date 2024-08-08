package hashes

// Define the length of a SKEIN-512 hash in hexadecimal form
const skein512Length = 128

// IsSKEIN512 checks if the hash is a valid SKEIN-512 hash
func IsSKEIN512(hash string) bool {
    return len(hash) == skein512Length && IsHex(hash)
}
