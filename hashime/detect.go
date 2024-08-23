// File: hashime/detect.go
package hashime

import (
	"github.com/0above/Hashime/hashime/hashes"
	"github.com/0above/Hashime/utils"
)

// HashType represents different types of hashes.
type HashType int

const (
	Unknown HashType = iota
	MD5
	SHA1
	SHA256
	SHA384
	SHA512
	SHA3_224
	SHA3_256
	SHA3_384
	SHA3_512
	RIPEMD160
	Blake2b
	CRC32
	CRC64
	FNV1
	FNV1a
	MurmurHash3
	BLAKE3
	Poly1305
	SKEIN512
	// Add other hash types here
)

// DetectHashType detects the type of hash based on the hash value.
func DetectHashType(hash string) HashType {
	if utils.IsHex(hash) {
		switch {
		case hashes.IsMD5(hash):
			return MD5
		case hashes.IsSHA1(hash):
			return SHA1
		case hashes.IsSHA256(hash):
			return SHA256
		case hashes.IsSHA384(hash):
			return SHA384
		case hashes.IsSHA512(hash):
			return SHA512
		case hashes.IsSHA3_224(hash):
			return SHA3_224
		case hashes.IsSHA3_256(hash):
			return SHA3_256
		case hashes.IsSHA3_384(hash):
			return SHA3_384
		case hashes.IsSHA3_512(hash):
			return SHA3_512
		case hashes.IsRIPEMD160(hash):
			return RIPEMD160
		case hashes.IsBlake2b(hash):
			return Blake2b
		// Add other cases here
		}
	}
	return Unknown
}

// HashTypeToString returns the string representation of HashType.
func HashTypeToString(ht HashType) string {
	switch ht {
	case MD5:
		return "MD5"
	case SHA1:
		return "SHA1"
	case SHA256:
		return "SHA256"
	case SHA384:
		return "SHA384"
	case SHA512:
		return "SHA512"
	case SHA3_224:
		return "SHA3-224"
	case SHA3_256:
		return "SHA3-256"
	case SHA3_384:
		return "SHA3-384"
	case SHA3_512:
		return "SHA3-512"
	case RIPEMD160:
		return "RIPEMD160"
	case Blake2b:
		return "Blake2b"
	case CRC32:
		return "CRC32"
	case CRC64:
		return "CRC64"
	case FNV1:
		return "FNV1"
	case FNV1a:
		return "FNV1a"
	case MurmurHash3:
		return "MurmurHash3"
	case BLAKE3:
		return "BLAKE3"
	case Poly1305:
		return "Poly1305"
	case SKEIN512:
		return "SKEIN512"
	default:
		return "Unknown"
	}
}
