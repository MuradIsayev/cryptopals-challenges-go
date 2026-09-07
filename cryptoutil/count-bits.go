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
