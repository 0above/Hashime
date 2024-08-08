package hashes

// Define the lengths of RIPEMD hashes in hexadecimal form
const (
    ripemd128Length = 32
    ripemd160Length = 40
    ripemd256Length = 64
    ripemd320Length = 80
)

// IsRIPEMD128 checks if the hash is a valid RIPEMD-128 hash
func IsRIPEMD128(hash string) bool {
    return len(hash) == ripemd128Length && IsHex(hash)
}

// IsRIPEMD160 checks if the hash is a valid RIPEMD-160 hash
func IsRIPEMD160(hash string) bool {
    return len(hash) == ripemd160Length && IsHex(hash)
}

// IsRIPEMD256 checks if the hash is a valid RIPEMD-256 hash
func IsRIPEMD256(hash string) bool {
    return len(hash) == ripemd256Length && IsHex(hash)
}

// IsRIPEMD320 checks if the hash is a valid RIPEMD-320 hash
func IsRIPEMD320(hash string) bool {
    return len(hash) == ripemd320Length && IsHex(hash)
}
