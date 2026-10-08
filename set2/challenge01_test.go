package set2

import (
	"crypto-challenges/cryptoutil"
	"testing"
)

func TestChallenge01_ImplementPKCS7_padding(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")
	expectedResult := "YELLOW SUBMARINE\x04\x04\x04\x04"

	gotResult, err := cryptoutil.PKCS7Padding(input, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if expectedResult != string(gotResult) {
		t.Fatalf("expected %q, got %q", expectedResult, gotResult)
	}

}
