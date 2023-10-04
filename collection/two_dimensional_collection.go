package collection

type TwoDimensionalCollection[K comparable, T any] struct {
	Collection[K, T]
	Items [][]T
}

func Of2D[K comparable, T any](d [][]T) *TwoDimensionalCollection[K, T] {
	return NewTwoDimensionalCollection[K, T](d)
}

func NewTwoDimensionalCollection[K comparable, T any](items [][]T) *TwoDimensionalCollection[K, T] {
	return &TwoDimensionalCollection[K, T]{
		Items: items,
	}
}

func (c *TwoDimensionalCollection[K, T]) First() []T {
	var empty []T

	if len(c.Items) > 0 {
		return c.Items[0]
	}

	return empty
}

func (c *TwoDimensionalCollection[K, T]) All() interface{} {
	return c.Items
}

func (c *TwoDimensionalCollection[K, T]) Values() [][]T {
	return c.Items
}
