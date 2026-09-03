package set1

import (
	"bufio"
	"crypto-challenges/cryptoutil"
	"fmt"
	"log"
	"math"
	"os"
)

func SingleCharXorInFile() {
	file, err := os.Open("./set1/testdata/4.txt")
	if err != nil {
		log.Fatal(err)
	}

	bestCandidateInFile := Candidate{
		Score: math.Inf(-1),
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bestCandidateByLine := SingleByteXorCipherByInput(scanner.Bytes())

		if bestCandidateByLine.Score > bestCandidateInFile.Score {
			bestCandidateInFile = Candidate{
				Key:       bestCandidateByLine.Key,
				Score:     bestCandidateByLine.Score,
				PlainText: bestCandidateByLine.PlainText,
			}
		}

		if !scanner.Scan() {
			break
		}
	}

	fmt.Printf(
		"\nSet 1.4 Answer: Key: %c (%d)\nScore: %.2f\nPlaintext: %s",
		bestCandidateInFile.Key,
		bestCandidateInFile.Key,
		bestCandidateInFile.Score,
		bestCandidateInFile.PlainText,
	)

}

func SingleByteXorCipherByInput(input []byte) *Candidate {
	decodedHex, err := cryptoutil.DecodeHex(input)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return nil
	}

	bestCandidate := FindBestCandidate(decodedHex)

	return &bestCandidate
}
