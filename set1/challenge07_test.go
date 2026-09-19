package set1

import (
	"crypto-challenges/cryptoutil"
	"crypto/aes"
	"io"
	"os"
	"testing"
)

func TestChallenge07_AES_in_ECB_mode(t *testing.T) {
	file, err := os.Open("testdata/7.txt")
	if err != nil {
		t.Fatalf("failed to open testdata/7.txt: %v", err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("failed to read content of testdata/7.txt: %v", err)
	}

	defer file.Close()

	decodedContent, err := cryptoutil.DecodeBase64(content)
	if err != nil {
		t.Fatalf("unexpected error decoding base64: %v", err)
	}

	block, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Fatalf("unexpected error creating a new cipher: %v", err)
	}

	dst := make([]byte, len(decodedContent))
	size := block.BlockSize()

	for i := 0; i < len(decodedContent); i += size {
		block.Decrypt(dst[i:i+size], decodedContent[i:i+size])
	}

	// Since the result is a full lyrics of a song, I'm leaving as it is.
	// t.Logf("result %v", string(dst))

}
