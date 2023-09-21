package collection

import (
	"reflect"
)

type OneDimensionalCollectionI[T any] interface {
	Collection[T]
	CollectFirst() T
	CollectLast() T
	CollectAll() []T
	Values() []T
}

type OneDimensionalCollection[T any] struct {
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

func (c *OneDimensionalCollection[T]) Each(fn func(T, int)) Collection[T] {
	for i, item := range c.Items {
		fn(item, i)
	}

	return c
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

func (c *OneDimensionalCollection[T]) collapse() *OneDimensionalCollection[T] {
	return c
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

func (c *OneDimensionalCollection[T]) Contains(fn func(T, int) bool) bool {
	flag := false

	for i, item := range c.Items {
		if fn(item, i) {
			flag = true
			break
		}
	}

	return flag
}

func (c *OneDimensionalCollection[T]) DoesntContain(fn func(T, int) bool) bool {
	flag := false

	for i, item := range c.Items {
		if fn(item, i) {
			flag = true
			break
		}
	}

	return !flag
}

func (c *OneDimensionalCollection[T]) Diff(elements []T) Collection[T] {
	var items []T
	el := make(map[any]bool)

	for _, v := range elements {
		el[v] = true
	}

	for _, v := range c.Items {
		if !el[v] {
			items = append(items, v)
		}
	}

	return &OneDimensionalCollection[T]{Items: items}
}

// Move this to maps collection
func (c *OneDimensionalCollection[T]) DiffAssoc(elements map[any]any) map[any]T {
	items := make(map[any]T)

	for i, v := range c.Items {
		if key, ok := elements[i]; ok {
			if !reflect.DeepEqual(key, v) {
				items[i] = v
			}
		}

	}

	return items
}

func (c *OneDimensionalCollection[T]) Duplicates() []T {
	var items []T
	arr := map[any]int{}

	for _, v := range c.Items {
		arr[v]++
	}

	for i, v := range arr {
		if v > 1 {
			items = append(items, i.(T))
		}
	}

	return items
}

func (c *OneDimensionalCollection[T]) DuplicatesBy(fieldName string) []any {
	var items []any
	arr := map[any]int{}

	for _, v := range c.Items {
		val := reflect.ValueOf(v)
		field := val.FieldByName(fieldName)
		if !field.IsValid() {
			panic("Field not found")
		}
		arr[field.Interface()]++
	}

	for i, v := range arr {
		if v > 1 {
			items = append(items, i.(any))
		}
	}

	return items
}

func (c *OneDimensionalCollection[T]) CollectFirst() T {
	var empty T

	if len(c.Items) > 0 {
		return c.Items[0]
	}

	return empty
}

func (c *OneDimensionalCollection[T]) First() interface{} {
	return c.CollectFirst()
}

func (c *OneDimensionalCollection[T]) CollectLast() T {
	var empty T

	if len(c.Items) > 0 {
		return c.Items[len(c.Items)-1]
	}

	return empty
}

func (c *OneDimensionalCollection[T]) Last() interface{} {
	return c.CollectLast()
}

func (c *OneDimensionalCollection[T]) Count() int {
	return len(c.Items)
}

func (c *OneDimensionalCollection[T]) CountBy(fn func(T, int) any) map[any]int {
	items := map[any]int{}
	if fn == nil {
		for _, v := range c.Items {
			items[v] += 1
		}
		return items
	}

	for i, v := range c.Items {
		key := fn(v, i)
		items[key] += 1
	}

	return items
}

func (c *OneDimensionalCollection[T]) CollectAll() []T {
	return c.Items
}

func (c *OneDimensionalCollection[T]) All() interface{} {
	return c.Items
}

func (c *OneDimensionalCollection[T]) Values() []T {
	return c.Items
}
