package collection

type Collection[T any] interface {
	Filter(fn func(T, int) bool) Collection[T]
	All() interface{}
	Values() Collection[T]
	Reject(fn func(T, int) bool) Collection[T]
}

func Of[T any](d interface{}) Collection[T] {
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
