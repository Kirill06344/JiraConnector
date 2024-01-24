package utils

func Filter[T any](arr []T, predicate func(T) bool) (ret []T) {
	for _, s := range arr {
		if predicate(s) {
			ret = append(ret, s)
		}
	}
	return
}
