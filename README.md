# Cryptopals Solutions in Go

[![Go Tests](https://github.com/MuradIsayev/cryptopals-challenges-go/actions/workflows/test.yml/badge.svg)](https://github.com/MuradIsayev/cryptopals-challenges-go/actions/workflows/test.yml)

My test-driven and straightforward Go implementations of the [Cryptopals Crypto Challenges](https://cryptopals.com/).

## Architecture & Approach
Instead of writing isolated scripts for each challenge, this repository is built as a cohesive, test-driven library. 

* **`cryptoutil/`**: A standalone package containing reusable cryptographic primitives (XOR logic, base64/hex encoding, English frequency scoring). This prevents code duplication as the challenges scale in complexity.
* **`set1/`**: The challenges themselves, implemented purely as Go table-driven tests (`_test.go`). This allows the entire suite to be verified instantly via standard Go tooling.

## Key Learnings: Set 1 (Basics)

* **Challenges 1 & 2 (Encodings & Fixed XOR):** Cryptography requires operating on raw bytes, not strings. Converting hex to base64 reinforces how data is packed (hex uses 4 bits per character; base64 uses 6 bits). Challenge 2 demonstrates the foundational, reversible property of XOR: $A \oplus B = C$, and $C \oplus B = A$.
* **Challenges 3 & 4 (Single-Byte XOR & Frequency Analysis):** Encryption is vulnerable if the underlying data has predictable patterns. By building a scoring function based on standard English letter frequencies (where 'e', 't', and ' ' appear most often), the plaintext mathematically bubbles to the top of hundreds of garbage decryptions. 
* **Challenge 5 (Repeating-Key XOR):** Transitioning from a single byte to a rotating key is handled cleanly using modular arithmetic (`key[i % len(key)]`). From a Go engineering perspective, this challenge highlighted the performance difference between continuously appending to a slice versus pre-allocating memory (`make([]byte, len(input))`) and assigning values by index.

## Running the Suite
To verify all challenge solutions, run the test suite from the root directory:

```bash
go test -v ./set1
