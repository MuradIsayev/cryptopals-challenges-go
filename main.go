package main

import "crypto-challenges/set1"

func main() {
	set1.HexToBase64()
	set1.FixedXor()
	set1.SingleByteXorCipher()
	set1.SingleCharXorInFile()

	input := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	key := []byte("ICE")
	set1.EncryptWithRepeatingKey(key, input)
}
