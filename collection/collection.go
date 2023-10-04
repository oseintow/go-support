package collection

type Collection[T any] interface {
	Filter(fn func(T, int) bool) Collection[T]
	Map(fn func(T, int) T) Collection[T]
	Each(fn func(T, int)) Collection[T]
	All() interface{}
	Reject(fn func(T, int) bool) Collection[T]
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
	Diff([]T) Collection[T]
	DiffAssoc(map[any]any) map[any]T
	Duplicates() []T
	DuplicatesBy(string) []any
	Every(fn func(T, int) bool) bool
	FirstWhere(fn func(T, int) bool) interface{}
	//except()
	FlatMap(fn func(T, int) []T) []T
	FlatMapAny(fn func(T, int) []any) []any
}

func Of[T any](d interface{}) Collection[T] {
	return NewCollection[T](d)
}

func NewCollection[T any](d interface{}) Collection[T] {
	switch r := d.(type) {
	case []T:
		return Of1D(r)
	case [][]T:
		//return Of2D(r)
		panic("two Dimensional collection is a work in progress")
	default:
		panic("unsupported type")
	}
}
