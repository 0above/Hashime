package hashes

// Define the lengths of CRC hashes in hexadecimal form
const (
    crc32Length = 8
    crc64Length = 16
)

// IsCRC32 checks if the hash is a valid CRC32 hash
func IsCRC32(hash string) bool {
    return len(hash) == crc32Length && IsHex(hash)
}

// IsCRC64 checks if the hash is a valid CRC64 hash
func IsCRC64(hash string) bool {
    return len(hash) == crc64Length && IsHex(hash)
}
