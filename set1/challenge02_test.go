package set1

import (
	"crypto-challenges/cryptoutil"
	"testing"
)

func TestChallenge02_FixedXOR(t *testing.T) {
	input1 := []byte("1c0111001f010100061a024b53535009181c")
	input2 := []byte("686974207468652062756c6c277320657965")
	expectedHex := "746865206b696420646f6e277420706c6179"

	decodedHex1, err := cryptoutil.DecodeHex(input1)
	if err != nil {
		t.Fatalf("unexpected error decoding hex: %v", err)
	}

	decodedHex2, err := cryptoutil.DecodeHex(input2)
	if err != nil {
		t.Fatalf("unexpected error decoding hex: %v", err)
	}

	result, err := cryptoutil.FixedXOR(decodedHex1, decodedHex2)
	if err != nil {
		t.Fatalf("unexpected error applying fixed XOR: %v", err)
	}

	gotHex := cryptoutil.EncodeHex(result)

	if string(gotHex) != expectedHex {
		t.Fatalf("expected %q, got %q", expectedHex, gotHex)
	}
}
