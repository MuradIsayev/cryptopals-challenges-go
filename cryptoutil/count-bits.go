package cryptoutil

func CountBitsKernighan(b byte) int {
	var count int
	for b > 0 {
		b &= (b - 1)
		count++
	}

	return count
}

func CountOnesInByte(b byte) int {
	var count int

	for b > 0 {
		count += int(b & 1)
		b >>= 1
	}

	return count
}

func CountOnesInString(b []byte) int {
	var count int

	for i := range b {
		currentByte := b[i]
		for currentByte > 0 {
			count += int(currentByte & 1)
			currentByte >>= 1
		}
	}

	return count
}

func FindHammingDistance(a, b []byte) (int, error) {
	xorResult, err := FixedXOR(a, b)
	if err != nil {
		return 0, err
	}

	var editDistance int
	editDistance = CountOnesInString(xorResult)

	return editDistance, nil
}
