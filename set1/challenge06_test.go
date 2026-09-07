package set1

import (
	"crypto-challenges/cryptoutil"
	"os"
	"testing"
)

func TestChallenge06_BreakRepeatingKeyXOR(t *testing.T) {
	file, err := os.Open("testdata/6.txt")
	input := 29
	expectedCount := 4
	if err != nil {
		t.Fatalf("failed to open testdata/4.txt: %v", err)
	}

	defer file.Close()

	getCount := cryptoutil.CountOnesInByte(byte(input))

	if getCount != expectedCount {
		t.Fatalf("expected %v, got %v", expectedCount, getCount)
	}
}
