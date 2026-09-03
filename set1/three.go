package set1

import (
	"crypto-challenges/cryptoutil"
	"fmt"
	"math"
)

type Candidate struct {
	Key       byte
	Score     float64
	PlainText string
}

func SingleByteXorCipher() {
	input := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	decodedHex, err := decodeHex(input)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return
	}

	bestCandidate := FindBestCandidate(decodedHex)

	fmt.Printf(
		"Key: %c (%d)\nScore: %.2f\nPlaintext: %s\n",
		bestCandidate.Key,
		bestCandidate.Key,
		bestCandidate.Score,
		bestCandidate.PlainText,
	)
}

func FindBestCandidate(cipherText []byte) Candidate {
	bestCandidate := Candidate{
		Score: math.Inf(-1),
	}

	for key := range 256 {
		plainText := SingleByteXOR(cipherText, byte(key))
		score := cryptoutil.CalculateEnglishScore(plainText)

		if score > bestCandidate.Score {
			bestCandidate = Candidate{
				Key:       byte(key),
				Score:     score,
				PlainText: string(plainText),
			}
		}

	}

	return bestCandidate
}

func SingleByteXOR(cipherText []byte, key byte) []byte {
	plainText := make([]byte, len(cipherText))

	for i, v := range cipherText {
		plainText[i] = v ^ key
	}

	return plainText
}
