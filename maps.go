package lib

func MapKeys[K comparable, V any](m map[K]V) (keys []K) {
	keys = make([]K, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	return keys
}
