# ![Hashime Logo](link-to-your-logo.png) Hashime

**Hashime** is your go-to library for detecting and working with various hash types in Go. Whether you’re building security tools or just need to identify a hash, Hashime makes it easy and efficient.

[![Go Report Card](https://goreportcard.com/badge/github.com/0above/Hashime)](https://goreportcard.com/report/github.com/0above/Hashime)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/0above/Hashime?status.svg)](https://godoc.org/github.com/0above/Hashime)

## 🎯 What is Hashime?

Hashime is a lightweight Go library designed to help you detect and identify different types of cryptographic hashes. From MD5 to SHA-512, we’ve got you covered.

## 🚀 Getting Started

Ready to dive in? Here’s how to get started with Hashime.

### Installation

Getting Hashime up and running is a breeze. Simply add it to your project with:

```bash
go get github.com/0above/Hashime/hashime
```

## Usage

Here’s a quick example of how to use Hashime for detecting hash types:

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/0above/Hashime/hashime"
)

func main() {
    hash := "5d41402abc4b2a76b9719d911017c592" // Example MD5 hash

    // Detect the hash type
    hashType := hashime.DetectHashType(hash)

    // Convert the detected hash type to string and print it
    fmt.Printf("Detected Hash Type: %s\n", hashime.HashTypeToString(hashType))
}
```

### Terminal Example

```go
package main

import (
    "fmt"
    "os"
    "github.com/0above/Hashime/hashime"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a hash value.")
        return
    }

    hash := os.Args[1]

    // Detect the hash type
    hashType := hashime.DetectHashType(hash)

    // Convert the detected hash type to string and print it
    fmt.Printf("Detected Hash Type: %s\n", hashime.HashTypeToString(hashType))
}
```

Run it from your terminal:

```sh
go run main.go 5d41402abc4b2a76b9719d911017c592
```

## ✨ Features

- **Wide Hash Support**: Detects various hash types including MD5, SHA1, SHA256, SHA512, and more.
- **Simple API**: Easy-to-use functions for detecting hash types.
- **Extensible**: Easily add support for new hash types as needed.
- **Lightweight**: Minimal dependencies, keeping your projects fast and efficient.

### Advanced Usage

For more advanced usage and configuration, refer to the [examples](https://github.com/0above/Hashime/tree/main/examples) directory of the repository.

## Examples

You can find more examples in the [examples directory](https://github.com/0above/Hashime/tree/main/examples) of the repository.

## Contributing

Contributions are welcome! Please read our [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute to the project.

## License

Hashime is released under the MIT License. See [LICENSE](LICENSE) for more details.

## Contact

For questions or feedback, please open an issue on the [GitHub repository](https://github.com/0above/Hashime) or contact [email@example.com](mailto:email@example.com).
