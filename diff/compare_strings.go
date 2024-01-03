package diff

func CompareStrings(s1, s2 string, pad int) (s string) {
	return newPair(s1, s2, pad).compareStrings()
}
