package lib

// LeadingSpaces returns the number of leading spaces in a string, e.g.
//
//	LeadingSpaces("foo")    => 0
//	LeadingSpaces(" foo")   => 1
//	LeadingSpaces("  foo")  => 2
//	LeadingSpaces("   foo") => 3
func LeadingSpaces(s string) (n int) {
	if s == "" {
		goto end
	}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			goto end
		}
		n++
	}
end:
	return n
}

// TrailingSpaces returns the number of trailing spaces in a string, e.g.
//
//	TrailingSpaces("bar")    => 0
//	TrailingSpaces("bar ")   => 1
//	TrailingSpaces("bar  ")  => 2
//	TrailingSpaces("bar   ") => 3
func TrailingSpaces(s string) (n int) {
	if s == "" {
		goto end
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			goto end
		}
		n++
	}
end:
	return n
}
