package set

type Set[K comparable, V any] interface {
	~map[K]V
}

func Empty[K comparable, S Set[K, struct{}]]() S {
	return make(map[K]struct{})
}

func Keys[K comparable, V any, S Set[K, V]](set S) []K {
	keys := make([]K, len(set))
	i := 0
	for k := range set {
		keys[i] = k
		i++
	}
	return keys
}

func Union[K comparable, V any, S Set[K, V]](sets ...S) map[K]struct{} {
	union := Empty[K]()
	for _, set := range sets {
		for key := range set {
			union[key] = struct{}{}
		}
	}
	return union
}

func Intersection[K comparable, V any, S Set[K, V]](sets ...S) map[K]struct{} {
	inter := Empty[K]()

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
