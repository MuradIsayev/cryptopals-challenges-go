package set1

import (
	"fmt"
)

func EncryptWithRepeatingKey(key, input []byte) {
	var result []byte
	// result := make([]byte, len(input))
	for i, val := range input {
		result = append(result, val^(key[i%len(key)]))
		// result[i] = val ^ (key[i%len(key)])
	}

	hexEncodedResult := encodeHex(result)

	fmt.Printf("Set One Five Answer: %s", hexEncodedResult)
}
