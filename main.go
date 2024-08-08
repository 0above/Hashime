package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/0above/Hashime/hashime" // Import hashime package
)


func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter hash: ")
    hash, _ := reader.ReadString('\n')
    hash = hash[:len(hash)-1] // Remove newline character

    hashType := hashime.DetectHashType(hash)
    fmt.Printf("Detected hash type: %v\n", hashType)
}
