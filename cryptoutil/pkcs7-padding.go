package cryptoutil

import (
	"fmt"
)

func PKCS7Padding(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 || blockSize > 255 {
		return nil, fmt.Errorf("invalid block size %d: PKCS#7 only supports 1 to 255", blockSize)
	}
	paddingLength := blockSize - (len(data) % blockSize)

	// padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)

	padded := make([]byte, len(data)+paddingLength)
	copy(padded, data)

	for i := len(data); i < len(padded); i++ {
		padded[i] = byte(paddingLength)
	}

	return padded, nil

}
