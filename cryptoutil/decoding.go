package cryptoutil

import "encoding/hex"

func DecodeHex(input []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(input))) // 2 hex characters decode to 1 byte
	// var byteCount int
	_, err := hex.Decode(dst, input)
	if err != nil {
		return nil, err
	}

	return dst, nil
}
