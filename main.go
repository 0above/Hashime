// File: main.go
package main

import (
	"flag"
	"fmt"

	"github.com/0above/Hashime/hashime" // Import the hashime package
)

func main() {
	hash := flag.String("hash", "", "Hash value to detect")
	flag.Parse()

	if *hash == "" {
		fmt.Println("Please provide a hash value with the -hash flag.")
		return
	}

	hashType := hashime.DetectHashType(*hash) // Call DetectHashType
	fmt.Printf("Detected hash type: %s\n", hashime.HashTypeToString(hashType)) // Call HashTypeToString
}
