package main

import (
	"crypto-challanges/setOne"
)

func main() {
	// setOne.HexToBase64()
	// setOne.FixedXor()
	// setOne.SingleByteXorCipher()
	// setOne.SingleCharXorInFile()

	input := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	key := []byte("ICE")
	setOne.EncryptWithRepeatingKey(key, input)
}
