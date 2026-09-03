package cryptoutil

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
