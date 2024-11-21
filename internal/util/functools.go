package util

func Map[T, V comparable](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func GetOrNil[T, V comparable](m map[T]V, key T) *V {
	val, ok := m[key]
	if ok {
		return &val
	}
	return nil
}

func Keys[T, V comparable](m map[T]V) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Filter[T comparable](ts []T, fn func(T) bool) []T {
	res := []T{}
	for _, t := range ts {
		if fn(t) {
			res = append(res, t)
		}
	}
	return res
}

func Replace[T comparable](ts []T, orig T, repl T) []T {
	res := []T{}
	for _, t := range ts {
		if t == orig {
			res = append(res, repl)
		} else {
			res = append(res, repl)
		}
	}
	return res
}

func Contains[T comparable](ts []T, elem T) bool {
	for _, t := range ts {
		if t == elem {
			return true
		}
	}
	return false
}

func Sum[T int64 | float64 | int | int32 | float32](ts []T) (o T) {
	var res T
	for _, pred := range ts {
		res += pred
	}
	return res
}

func Any(ts []bool) bool {
	for _, t := range ts {
		if t {
			return true
		}
	}
	return false
}

// Represents given slice as a set. By convention, a set
// is a [map] with keys of type [T] and values of the
// [struct{}] type.
func SliceToSet[T comparable](ts []T) (res map[T]any) {
	res = map[T]any{}
	for _, e := range ts {
		res[e] = struct{}{}
	}
	return res
}

func SetToSlice[T comparable](ts map[T]any) (res []T) {
	res = []T{}
	for k := range ts {
		res = append(res, k)
	}
	return res
}

func Unique[T comparable](ts []T) []T {
	return SetToSlice(SliceToSet(ts))
}

// This function intersects two sets, represented
// as [map]s with keys of type [T] and values of
// [any] type (ideally, [struct{}]).
func Intersect[T comparable, V any](s1, s2 map[T]V) (res map[T]any) {
	res = map[T]any{}
	for e1 := range s1 {
		for e2 := range s2 {
			if e1 == e2 {
				res[e1] = struct{}{}
			}
		}
	}
	return res
}
