package lib

func RightShift[S []E, E any](s S, e ...E) S {
	s = append(s, e...)
	if len(s) != len(e) {
		copy(s[len(e):], s[:len(s)-len(e)])
		copy(s, e)
	}
	return s
}
