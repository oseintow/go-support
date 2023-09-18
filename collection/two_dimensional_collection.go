package collection

type TwoDimensionalCollection[T any] struct {
	Collection[T]
	Items [][]T
}

func Of2D[T any](d [][]T) *TwoDimensionalCollection[T] {
	return NewTwoDimensionalCollection(d)
}

func NewTwoDimensionalCollection[T any](items [][]T) *TwoDimensionalCollection[T] {
	return &TwoDimensionalCollection[T]{
		Items: items,
	}
}

func (c *TwoDimensionalCollection[T]) First() []T {
	var empty []T

	if len(c.Items) > 0 {
		return c.Items[0]
	}

	return empty
}

func (c *TwoDimensionalCollection[T]) All() interface{} {
	return c.Items
}

func (c *TwoDimensionalCollection[T]) Values() [][]T {
	return c.Items
}
