package set1

import (
	"crypto-challenges/cryptoutil"
	"fmt"
)

func HexToBase64() {
	hexInput := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	decodedHex, err := cryptoutil.DecodeHex(hexInput)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return
	}

	encodedBase64, err := cryptoutil.EncodeBase64(decodedHex)
	if err != nil {
		fmt.Printf("failed to encode to base64: %s", err)
		return
	}

	fmt.Printf("\nSet 1.1 Answer: %s", encodedBase64)
}
