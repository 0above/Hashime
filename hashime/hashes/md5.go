package hashes

// Define the length of an MD5 hash in hexadecimal form
const md5Length = 32

// IsMD5 checks if the hash is a valid MD5 hash
func IsMD5(hash string) bool {
    return len(hash) == md5Length && IsHex(hash)
}
