package collection

type Collection[T any] struct {
	Items []T
}

func Of[T any](d []T) *Collection[T] {
	return NewCollection(d)
}

func NewCollection[T any](items []T) *Collection[T] {
	return &Collection[T]{
		Items: items,
	}
}

func (c *Collection[T]) Filter(filterFunc func(T, int) bool) *Collection[T] {
	var items []T

	for i, item := range c.Items {
		if filterFunc(item, i) {
			items = append(items, item)
		}
	}

	return &Collection[T]{items}
}

func (c *Collection[T]) Reject(rejectFunc func(T, int) bool) *Collection[T] {
	var items []T

	for i, v := range c.Items {
		if !rejectFunc(v, i) {
			items = append(items, v)
		}
	}

	return &Collection[T]{items}
}

func (c *Collection[T]) Map(mapFunc func(T, int) T) *Collection[T] {
	var items []T

	for i, item := range c.Items {
		items = append(items, mapFunc(item, i))
	}

	return &Collection[T]{items}
}

func (c *Collection[T]) Chunk(size int) [][]T {
	var items [][]T

	for i := 0; i < len(c.Items); i += size {
		j := i + size
		if j > len(c.Items) {
			j = len(c.Items)
		}
		items = append(items, c.Items[i:j])
	}

	return items
}

func (c *Collection[T]) First() T {
	var empty T

	if len(c.Items) > 0 {
		return c.Items[0]
	}

	return empty
}

func (c *Collection[T]) Last() T {
	var empty T

	if len(c.Items) > 0 {
		return c.Items[len(c.Items)-1]
	}

	return empty
}

func (c *Collection[T]) Count() int {
	return len(c.Items)
}

func (c *Collection[T]) All() []T {
	return c.Items
}

func (c *Collection[T]) Values() *Collection[T] {
	return c
}
