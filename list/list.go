package list

func Reverse[T any](list []T) []T {
	n := len(list)
	reversed := make([]T, n)
	for i, v := range list {
		reversed[n-i-1] = v
	}
	return reversed
}
