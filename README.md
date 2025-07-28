![Hashime Logo](assets/gopher_sushi1.png)

# Hello, World!

**Welcome to Hashime** — a pure-Go hash detection library crafted for Go developers.

Hashime takes any hash string, analyzes it with multiple detection techniques, and returns the most likely hash type along with a confidence score. It’s fast, lightweight, and has no external dependencies.

---

## Project Status

When I first built Hashime, I was learning Go and got the core detection working. Life intervened, and development paused. Now, I’m back and committed to a full revamp. I’m rebuilding the codebase locally and will upload polished, complete updates rather than pushing every small commit. This way, you get stable, ready-to-use releases.

---

## The Revamp: What’s Changing

* **Simplified architecture**: A cleaner, more maintainable code structure.
* **Performance improvements**: Faster detection routines and reduced memory usage.
* **Modular design**: Components you can reuse and extend in your own projects.

---

## Upcoming Features

* **Salt detection**: Flag hashes that include a salt value.
* **Broader hash coverage**: Support for Streebog, LM Hash, and other algorithms.
* **Enhanced confidence scoring**: Smarter heuristics to improve accuracy.
* **Improved CLI experience**: Cleaner terminal output, optional color highlights, and a simple banner.

---

## Features (Today)

* **Library + CLI**: Import as a Go module or run a standalone command-line tool.
* **Pure Go**: No third-party libraries—just add it to your `go.mod`.
* **Multi-method detection**: Combines pattern checks, length analysis, and entropy measures for reliable results.

---

## Installation & Usage

* **This section describes the current (old) build.

### As a Go module

```bash
go get github.com/0above/Hashime
```

In your code:

```go
usage 

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
```

### From the terminal

```bash
go run main.go 5d41402abc4b2a76b9719d911017c592
```

Output example:

```
Detected: MD5
```

---

## Roadmap & Contributions

* Stable core: Reliable detection for MD5, SHA-1, SHA-256, and more.
* Active development: Complete rewrite in progress with exciting new features.

**Expected Releases:**

* **October 2025**: First usable revamp version with core new features integrated.
* **December 2025**: Fully revamped release with all planned enhancements completed.

Because I’m uploading finished updates only, GitHub will show stable snapshots of each revamp stage. To contribute:

1. Fork the repository
2. Create a feature branch
3. Submit a pull request or open an issue

License: [MIT](LICENSE)

---

## Feedback & Ideas

I’d love to hear from you:

* Which hash algorithms matter most to you?
* What CLI features would make Hashime indispensable?
* Any sample hashes or test cases we should include?

Thanks for checking out Hashime—let’s build something great together! 🎉
