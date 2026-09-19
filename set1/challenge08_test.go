package set1

import (
	"bufio"
	"crypto-challenges/cryptoutil"
	"os"
	"testing"
)

func TestChallenge08_Detect_AES_in_ECB_Mode(t *testing.T) {
	file, err := os.Open("testdata/8.txt")
	if err != nil {
		t.Fatalf("failed to open testdata/8.txt: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	blockSize := 16
	foundECB := false

	for scanner.Scan() {
		lineNumber++
		line := scanner.Bytes()

		decodedHex, err := cryptoutil.DecodeHex(line)
		if err != nil {
			t.Fatalf("unexpected error decoding hex: %v", err)
		}

		duplicateChunks := 0
		chunkFrequency := make(map[string]int)

		for i := 0; i+blockSize <= len(decodedHex); i += blockSize {
			chunk := string(decodedHex[i : i+blockSize])
			if chunkFrequency[chunk] > 0 {
				duplicateChunks++
			}

			chunkFrequency[chunk]++
		}

		if duplicateChunks > 0 {
			foundECB = true
			// t.Logf("Found ECB on Line %d containing %d duplicate blocks.", lineNumber, duplicateChunks)
			// t.Logf("Ciphertext: %s", line)
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	if !foundECB {
		t.Fatal("Failed to detect any ECB encrypted strings in the file")
	}
}
