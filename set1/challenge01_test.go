package set1

import (
	"crypto-challenges/cryptoutil"
	"testing"
)

func TestChallenge01_HextoBase64(t *testing.T) {
	input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	decodedHex, err := cryptoutil.DecodeHex(input)
	if err != nil {
		t.Fatalf("unexpected error decoding hex: %v", err)
	}

	gotBase64, err := cryptoutil.EncodeBase64(decodedHex)
	if err != nil {
		t.Fatalf("unexpected error encoding base64: %v", err)
	}

	if string(gotBase64) != expectedBase64 {
		t.Fatalf("expected %q, got %q", expectedBase64, gotBase64)
	}
}
