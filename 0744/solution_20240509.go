package _744

func nextGreatestLetter(letters []byte, target byte) byte {

	lo, hi := 0, len(letters)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if letters[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if letters[hi] <= target {
		return letters[0]
	}

	return letters[hi]
}
