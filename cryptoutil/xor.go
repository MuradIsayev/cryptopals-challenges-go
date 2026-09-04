package cryptoutil

import "fmt"

func RepeatingKeyXOR(input, key []byte) []byte {
	if len(key) == 0 {
		return input
	}

	out := make([]byte, len(input))
	// var out []byte
	for i, val := range input {
		// out = append(out, val^(key[i%len(key)]))
		out[i] = val ^ (key[i%len(key)])
	}

	return out
}

func FixedXOR(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("cryptoutil: buffer lengths must be equal (%d != %d)", len(a), len(b))
	}

	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}

	return out, nil
}

func SingleByteXOR(cipherInput []byte, key byte) []byte {
	decrpytedOut := make([]byte, len(cipherInput))

	for i, v := range cipherInput {
		decrpytedOut[i] = v ^ key
	}

	return decrpytedOut
}
