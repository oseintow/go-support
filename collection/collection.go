package collection

type Collection[K comparable, T any] interface {
	Map(fn func(T, int) T) Collection[K, T]
	Filter(fn func(T, int) bool) Collection[K, T]
	Each(fn func(T, int)) Collection[K, T]
	All() interface{}
	Reject(fn func(T, int) bool) Collection[K, T]
	Chunk(int) [][]T
	Average(fn func(T) float64) float64
	Avg(fn func(T) float64) float64
	Combine([]T) map[any]T
	Contains(fn func(T, int) bool) bool
	DoesntContain(fn func(T, int) bool) bool
	First() interface{}
	Last() interface{}
	Count() int
	CountBy(fn func(T, int) any) map[any]int
	Diff([]T) Collection[K, T]
	DiffAssoc(map[any]any) map[any]T
	Duplicates() []T
	DuplicatesBy(string) []any
	Every(fn func(T, int) bool) bool
	FirstWhere(fn func(T, int) bool) interface{}
	//except()
	FlatMap(fn func(T, int) []K) Collection[K, K]
}

func Of[K comparable, T any](items []T) Collection[K, T] {
	return newCollection[K, T](items)
}

func newCollection[K comparable, T any](items interface{}) Collection[K, T] {
	switch r := items.(type) {
	case []T:
		return Of1D[K, T](r)
	case [][]T:
		//return Of2D(r)
		panic("two Dimensional collection is a work in progress")
	default:
		panic("unsupported type")
	}
}
