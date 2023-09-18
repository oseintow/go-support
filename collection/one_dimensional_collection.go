package collection

type OneDimensionalCollectionI[T any] interface {
	First1D() T
	Last1D() T
	All1D() []T
	Values() []T
}

type OneDimensionalCollection[T any] struct {
	OneDimensionalCollectionI[T]
	Items []T
}

func Of1D[T any](d []T) *OneDimensionalCollection[T] {
	return NewOneDimensionalCollection(d)
}

func NewOneDimensionalCollection[T any](items []T) *OneDimensionalCollection[T] {
	return &OneDimensionalCollection[T]{
		Items: items,
	}
}

func (c *OneDimensionalCollection[T]) Filter(fn func(T, int) bool) Collection[T] {
	var items []T

	for i, item := range c.Items {
		if fn(item, i) {
			items = append(items, item)
		}
	}

	return &OneDimensionalCollection[T]{Items: items}
}

func (c *OneDimensionalCollection[T]) Reject(fn func(T, int) bool) Collection[T] {
	var items []T

	for i, v := range c.Items {
		if !fn(v, i) {
			items = append(items, v)
		}
	}

	return &OneDimensionalCollection[T]{Items: items}
}

func (c *OneDimensionalCollection[T]) Map(fn func(T, int) T) Collection[T] {
	var items []T

	for i, item := range c.Items {
		items = append(items, fn(item, i))
	}

	return &OneDimensionalCollection[T]{Items: items}
}

func (c *OneDimensionalCollection[T]) Chunk(size int) [][]T {
	if size < 1 {
		panic("size should be greater than zero(0)")
	}

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

// TODO: Collapse
func (c *OneDimensionalCollection[T]) collapse() {
	//var items []T

	//for _, item := range c.Items {
	//	items = append(items, item...)
	//}
}

func (c *OneDimensionalCollection[T]) Combine(elements []T) map[any]T {
	if len(c.Items) != len(elements) {
		panic("The two slices/arrays must have equal length")
	}

	items := make(map[any]T, len(c.Items))

	for i, v := range c.Items {
		items[v] = elements[i]
	}

	return items
}

func (c *OneDimensionalCollection[T]) Contains(containsFunc func(T, int) bool) bool {
	flag := false

	for i, item := range c.Items {
		if containsFunc(item, i) {
			flag = true
			break
		}
	}

	return flag
}

func (c *OneDimensionalCollection[T]) First1D() T {
	var empty T

	if len(c.Items) > 0 {
		return c.Items[0]
	}

	return empty
}

func (c *OneDimensionalCollection[T]) First() interface{} {
	return c.First1D()
}

func (c *OneDimensionalCollection[T]) Last1D() T {
	var empty T

	if len(c.Items) > 0 {
		return c.Items[len(c.Items)-1]
	}

	return empty
}

func (c *OneDimensionalCollection[T]) Last() interface{} {
	return c.Last1D()
}

func (c *OneDimensionalCollection[T]) Count() int {
	return len(c.Items)
}

func (c *OneDimensionalCollection[T]) CountBy(countByFunc func(T, int) any) map[any]int {
	items := map[any]int{}
	if countByFunc == nil {
		for _, v := range c.Items {
			items[v] += 1
		}
		return items
	}

	for i, v := range c.Items {
		key := countByFunc(v, i)
		items[key] += 1
	}

	return items
}

func (c *OneDimensionalCollection[T]) All1D() []T {
	return c.Items
}

func (c *OneDimensionalCollection[T]) All() interface{} {
	return c.Items
}

func (c *OneDimensionalCollection[T]) Values() []T {
	return c.Items
}
