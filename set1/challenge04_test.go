package set1

import (
	"bufio"
	"crypto-challenges/cryptoutil"
	"math"
	"os"
	"testing"
)

func TestChallenge04_DetectSingleCharXOR(t *testing.T) {
	expectedDecryptedOut := "Now that the party is jumping\n"

	file, err := os.Open("testdata/4.txt")
	if err != nil {
		t.Fatalf("failed to open testdata/4.txt: %v", err)
	}
	defer file.Close()

	bestCandidateInFile := cryptoutil.Candidate{
		Score: math.Inf(-1),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) == 0 {
			continue
		}

		decodedHex, err := cryptoutil.DecodeHex(line)
		if err != nil {
			t.Fatalf("unexpected error decoding hex on line %q: %v", string(line), err)
		}

		bestCandidateForLine := cryptoutil.FindBestCandidate(decodedHex)

		if bestCandidateForLine.Score > bestCandidateInFile.Score {
			bestCandidateInFile = bestCandidateForLine
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	if bestCandidateInFile.DecryptedOut != expectedDecryptedOut {
		t.Fatalf("expected %q, got %q (score: %.2f, key: %c)",
			expectedDecryptedOut,
			bestCandidateInFile.DecryptedOut,
			bestCandidateInFile.Score,
			bestCandidateInFile.Key,
		)
	}
}
