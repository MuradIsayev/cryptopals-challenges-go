package setOne

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func FixedXor() {
	input1 := []byte("1c0111001f010100061a024b53535009181c")
	input2 := []byte("686974207468652062756c6c277320657965")

	decodedInput1, err := decodeHex(input1)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return

	}

	decodedInput2, err := decodeHex(input2)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return

	}

	xorResult, err := xor(decodedInput1, decodedInput2)
	if err != nil {
		fmt.Printf("failed to xor inputs: %s", err)
		return
	}

	encodedHex := encodeHex(xorResult)

	fmt.Printf("Set One Two Answer: %s", encodedHex)
}

func encodeHex(input []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(input)))

	byteCount := hex.Encode(dst, input)
	fmt.Printf("%d bytes written\n", byteCount)

	return dst
}

func xor(a, b []byte) ([]byte, error) {
	inputLength := len(a)
	dst := make([]byte, inputLength)
	if notEqual := (len(a) != len(b)); notEqual {
		return nil, errors.New("Buffers have different lengths")
	}

	for i := range a {
		dst[i] = a[i] ^ b[i]
	}

	return dst, nil

}
