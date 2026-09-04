package set1

import (
	"crypto-challenges/cryptoutil"
	"testing"
)

// Single-byte XOR cipher

func TestChallenge03_SingleByteXORCipher(t *testing.T) {
	input := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	expectedDecryptedOut := "Cooking MC's like a pound of bacon"

	decodedHex, err := cryptoutil.DecodeHex(input)
	if err != nil {
		t.Fatalf("unexpected error decoding hex: %v", err)
	}

	bestCandidate := cryptoutil.FindBestCandidate(decodedHex)

	if bestCandidate.DecryptedOut != expectedDecryptedOut {
		t.Fatalf("expected %q, got %q (score: %.2f, key: %c)",
			expectedDecryptedOut,
			bestCandidate.DecryptedOut,
			bestCandidate.Score,
			bestCandidate.Key,
		)
	}

}
