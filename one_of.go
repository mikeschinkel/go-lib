package lib

func OneOf[T comparable](val T, options ...T) bool {
	for _, option := range options {
		if val == option {
			return true
		}
	}
	return false
}
