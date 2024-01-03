package diff

func CompareStrings(s1, s2 string, opts *CompareOpts) (s string) {
	opts.SetDefaults()
	return newPair(s1, s2, opts).compareStrings()
}
