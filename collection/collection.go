package collection

type Collection[T any] interface {
	Filter(fn func(T, int) bool) Collection[T]
	Map(fn func(T, int) T) Collection[T]
	All() interface{}
	Reject(fn func(T, int) bool) Collection[T]
	Chunk(size int) [][]T
	Average(avgFunc func(T) float64) float64
	Avg(avgFunc func(T) float64) float64
	Combine(elements []T) map[any]T
	Contains(containsFunc func(T, int) bool) bool
	First() interface{}
	Last() interface{}
	Count() int
	CountBy(countByFunc func(T, int) any) map[any]int
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
		panic("WIP: Two Dimensional Collection in progress")
	default:
		panic("unsupported type")
	}
}
