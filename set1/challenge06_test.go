package set1

import (
	"crypto-challenges/cryptoutil"
	"io"
	"math"
	"os"
	"testing"
)

func TestChallenge06_BreakRepeatingKeyXOR(t *testing.T) {
	expectedResult := "Terminator X: Bring the noise"
	file, err := os.Open("testdata/6.txt")
	if err != nil {
		t.Fatalf("failed to open testdata/6.txt: %v", err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("failed to read content of testdata/6.txt: %v", err)
	}
	defer file.Close()

	decodedContent, err := cryptoutil.DecodeBase64(content)
	if err != nil {
		t.Fatalf("unexpected error decoding base64: %v", err)
	}

	bestHamDistance := math.Inf(1)
	bestKeySize := 0

	for keySize := 2; keySize < 40; keySize++ {
		totalDistance := 0

		numBlocks := len(decodedContent) / keySize

		for i := 0; i < numBlocks-1; i++ {
			chunk1 := decodedContent[i*keySize : (i+1)*keySize]
			chunk2 := decodedContent[(i+1)*keySize : (i+2)*keySize]

			distance, err := cryptoutil.FindHammingDistance(chunk1, chunk2)
			if err != nil {
				t.Fatalf("unexpected error calculating distance: %v", err)
			}
			totalDistance += distance
		}

		normalizedHamDistance := float64(totalDistance) / (float64(numBlocks-1) * float64(keySize))

		if normalizedHamDistance < bestHamDistance {
			bestHamDistance = normalizedHamDistance
			bestKeySize = keySize
		}

	}
	// t.Logf("The winning KEYSIZE is: %d (Score: %.2f)", bestKeySize, bestHamDistance)

	transposedBlocks := make([][]byte, bestKeySize)
	for i, b := range decodedContent {
		bucketIndex := i % bestKeySize
		transposedBlocks[bucketIndex] = append(transposedBlocks[bucketIndex], b)
	}

	gotResult := make([]byte, 0, bestKeySize)
	for j := range bestKeySize {
		bestCandidate := cryptoutil.FindBestCandidate(transposedBlocks[j])
		gotResult = append(gotResult, bestCandidate.Key)
	}

	if expectedResult != string(gotResult) {
		t.Fatalf("expected %v, got %v", expectedResult, string(gotResult))
	}
}
