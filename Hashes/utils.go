package hashes

// IsHex checks if the string is a valid hexadecimal representation.
func IsHex(s string) bool {
    for _, c := range s {
        if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
            return false
        }
    }
    return true
}
