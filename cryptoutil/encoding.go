package cryptoutil

import (
	"encoding/base64"
	"encoding/hex"
)

func EncodeHex(input []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(input)))

	hex.Encode(dst, input)

	return dst
}

func EncodeBase64(input []byte) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(input)))

	base64.StdEncoding.Encode(dst, input)

	return dst, nil
}
