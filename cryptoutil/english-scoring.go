package cryptoutil

import "math"

// Candidate holds the results of a single-byte XOR decryption attempt.
type Candidate struct {
	Key          byte
	Score        float64
	DecryptedOut string
}

// Letter distributions in English text
var LetterFrequencies = map[byte]float64{
	'a': 8.167,
	'b': 1.492,
	'c': 2.782,
	'd': 4.253,
	'e': 12.702,
	'f': 2.228,
	'g': 2.015,
	'h': 6.094,
	'i': 6.966,
	'j': 0.153,
	'k': 0.772,
	'l': 4.025,
	'm': 2.406,
	'n': 6.749,
	'o': 7.507,
	'p': 1.929,
	'q': 0.095,
	'r': 5.987,
	's': 6.327,
	't': 9.056,
	'u': 2.758,
	'v': 0.978,
	'w': 2.360,
	'x': 0.150,
	'y': 1.974,
	'z': 0.074,
}

func CalculateEnglishScore(decryptedOut []byte) float64 {
	score := 0.0

	for _, c := range decryptedOut {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}

		switch {
		case c >= 'a' && c <= 'z':
			score += LetterFrequencies[c]

		case c == ' ':
			score += 5.0

		case c == '.', c == ',', c == '\'', c == '!', c == '?':
			score += 0.5

		case c >= 32 && c <= 126:
			score -= 1.0

		default:
			score -= 10.0
		}
	}

	return score

}

func FindBestCandidate(cipherInput []byte) Candidate {
	bestCandidate := Candidate{
		Score: math.Inf(-1),
	}

	for key := range 256 {
		decryptedOut := SingleByteXOR(cipherInput, byte(key))
		score := CalculateEnglishScore(decryptedOut)

		if score > bestCandidate.Score {
			bestCandidate = Candidate{
				Key:          byte(key),
				Score:        score,
				DecryptedOut: string(decryptedOut),
			}
		}
	}

	return bestCandidate
}
