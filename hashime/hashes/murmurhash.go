package hashes

// Define the length of a MurmurHash3 hash in hexadecimal form
const murmurHash3Length = 8

// IsMurmurHash3 checks if the hash is a valid MurmurHash3 hash
func IsMurmurHash3(hash string) bool {
    return len(hash) == murmurHash3Length && IsHex(hash)
}
