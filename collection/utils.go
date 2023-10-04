package collection

func FlatMap[K comparable, T any](collection []T, fn func(T, int) []K) []K {
	var items []K

	for i, v := range collection {
		items = append(items, fn(v, i)...)
	}

	return items
}

func Flatten[T any](collection [][]T) []T {
	var item []T

	for _, v := range collection {
		item = append(item, v...)
	}

	return item
}
