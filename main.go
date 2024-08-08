package main

import (
	"flag"
	"fmt"

	"github.com/0above/Hashime/hashime"
)

func main() {
    // Define and parse command-line flags
    var hash string
    flag.StringVar(&hash, "hash", "", "Hash value to detect")
    flag.Parse()

    // Show help if requested
    if flag.NFlag() == 0 {
        flag.Usage()
        return
    }

    // Ensure hash value is provided
    if hash == "" {
        fmt.Println("Please provide a hash value using the -hash flag.")
        flag.Usage()
        return
    }

    // Detect hash type
    hashType := hashime.DetectHashType(hash)
    switch hashType {
    case hashime.MD5:
        fmt.Println("MD5")
    case hashime.SHA1:
        fmt.Println("SHA1")
    case hashime.SHA256:
        fmt.Println("SHA256")
    case hashime.SHA512:
        fmt.Println("SHA512")
    default:
        fmt.Println("Unknown")
    }
}
