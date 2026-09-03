package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64() {
	hexInput := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	decodedHex, err := decodeHex(hexInput)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return
	}

	encodedBase64, err := encodeBase64(decodedHex)
	if err != nil {
		fmt.Printf("failed to encode to base64: %s", err)
		return
	}
	// fmt.Printf("%s\n", decodedHex)
	fmt.Printf("Set One One Answer: %s\n", encodedBase64)
}

func decodeHex(hexInput []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(hexInput))) // 2 hex characters decode to 1 byte
	// var byteCount int
	_, err := hex.Decode(dst, hexInput)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%d bytes written\n", byteCount)

	return dst, nil
}

func encodeBase64(input []byte) ([]byte, error) {
	dst := make([]byte, base64.RawStdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(dst, input)

	return dst, nil
}
