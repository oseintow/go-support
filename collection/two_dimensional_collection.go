package collection

type TwoDimensionalCollection[T any] struct {
	Items []T
}

func Of2D[T any](d []T) *TwoDimensionalCollection[T] {
	return NewTwoDimensionalCollection(d)
}

func NewTwoDimensionalCollection[T any](items []T) *TwoDimensionalCollection[T] {
	return &TwoDimensionalCollection[T]{
		Items: items,
	}
}
