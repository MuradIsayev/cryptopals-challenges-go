package cryptoutil

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

// func CountOnesKernighan(b byte) int {
// 	var count int
// 	for b > 0 {
// 		b &= (b - 1)
// 		count++
// 	}
//
// 	return count
// }
//
// func CountOnesIterativeShift(b byte) int {
// 	var count int
//
// 	for b > 0 {
// 		count += int(b & 1)
// 		b >>= 1
// 	}
//
// 	return count
// }
