# Hashime - A Go Library for Hash Detection

![Static Badge](https://img.shields.io/badge/:badgeContent?style=social)

## Overview

Hashime is a Go library designed for efficient hash detection and verification. It helps identify and verify different hash types, making it useful for building hash crackers, hashing your own code, and other hash-related tasks.

**Note: Hashime is currently in an early development stage (undersaturation). The library is not yet feature-complete, and changes are expected.**

![Hashime Logo](addlink)

## Features

- **Support for Multiple Hash Algorithms**: SHA-256, MD5, SHA-1, and more.
- **Efficient Hash Detection**: Designed for performance and ease of use.

## Installation

To install Hashime, use `go get`:

```sh
go get github.com/0above/Hashime
```

## Usage

Here’s a quick example of how to use Hashime for detecting hash types:

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/0above/Hashime"
)

func main() {
    data := []byte("example data")

    detector := hashime.New()

    sha256Hash := detector.CalculateSHA256(data)
    fmt.Printf("SHA-256: %x\n", sha256Hash)

    md5Hash := detector.CalculateMD5(data)
    fmt.Printf("MD5: %x\n", md5Hash)
}
```

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
