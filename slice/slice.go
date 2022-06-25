package slice

func Reverse[T any](list []T) []T {
	n := len(list)
	reversed := make([]T, n)
	for i, v := range list {
		reversed[n-i-1] = v
	}
	return reversed
}

func Includes[T comparable](slice []T, value T) bool {
	for _, it := range slice {
		if it == value {
			return true
		}
	}
	return false
}
