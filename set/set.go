package set

func Empty[K comparable, V any]() map[K]V {
	return make(map[K]V)
}

func Keys[K comparable, V any](s map[K]V) []K {
	keys := make([]K, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}

func Union[K comparable, V any](sets ...map[K]V) map[K]struct{} {
	union := Empty[K, struct{}]()
	for _, set := range sets {
		for key := range set {
			union[key] = struct{}{}
		}
	}
	return union
}

func Intersection[K comparable, V any](sets ...map[K]V) map[K]struct{} {
	inter := Empty[K, struct{}]()

	if len(sets) == 0 {
		return inter
	}
	if len(sets) == 1 {
		for key := range sets[0] {
			inter[key] = struct{}{}
		}
	}

	for key := range sets[0] {
		for _, b := range sets[1:] {
			if _, ok := b[key]; ok {
				inter[key] = struct{}{}
			}
		}
	}

	return inter
}
