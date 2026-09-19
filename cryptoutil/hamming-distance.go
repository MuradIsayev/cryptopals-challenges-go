package cryptoutil

func FindHammingDistance(a, b []byte) (int, error) {
	xorResult, err := FixedXOR(a, b)
	if err != nil {
		return 0, err
	}

	var editDistance int
	editDistance = CountOnesInString(xorResult)

	return editDistance, nil
}
