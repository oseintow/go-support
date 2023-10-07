package collection

import (
	"reflect"
)

type IOneDimensionalCollection[K comparable, T any] interface {
	CollectFirst() T
	CollectLast() T
	CollectAll() []T
	Values() []T
}

type OneDimensionalCollection[K comparable, T any] struct {
	Items []T
}

func Of1D[K comparable, T any](items []T) *OneDimensionalCollection[K, T] {
	//return NewOneDimensionalCollection[K, T](items)
	return &OneDimensionalCollection[K, T]{
		Items: items,
	}
}

func NewOneDimensionalCollection[K comparable, T any](items []T) *OneDimensionalCollection[K, T] {
	return &OneDimensionalCollection[K, T]{
		Items: items,
	}
}

func (c *OneDimensionalCollection[K, T]) Filter(fn func(T, int) bool) Collection[K, T] {
	var items []T

	for i, item := range c.Items {
		if fn(item, i) {
			items = append(items, item)
		}
	}

	return &OneDimensionalCollection[K, T]{Items: items}
}

func (c *OneDimensionalCollection[K, T]) Reject(fn func(T, int) bool) Collection[K, T] {
	var items []T

	for i, v := range c.Items {
		if !fn(v, i) {
			items = append(items, v)
		}
	}

	return &OneDimensionalCollection[K, T]{Items: items}
}

func (c *OneDimensionalCollection[K, T]) Map(fn func(T, int) T) Collection[K, T] {
	var items []T

	for i, item := range c.Items {
		items = append(items, fn(item, i))
	}

	return &OneDimensionalCollection[K, T]{Items: items}
}

func (c *OneDimensionalCollection[K, T]) Each(fn func(T, int)) Collection[K, T] {
	for i, item := range c.Items {
		fn(item, i)
	}

	return c
}

func (c *OneDimensionalCollection[K, T]) Chunk(size int) [][]T {
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

func (c *OneDimensionalCollection[K, T]) collapse() *OneDimensionalCollection[K, T] {
	return c
}

func (c *OneDimensionalCollection[K, T]) Combine(elements []T) map[any]T {
	if len(c.Items) != len(elements) {
		panic("The two slices/arrays must have equal length")
	}

	items := make(map[any]T, len(c.Items))

	for i, v := range c.Items {
		items[v] = elements[i]
	}

	return items
}

func (c *OneDimensionalCollection[K, T]) Contains(fn func(T, int) bool) bool {
	flag := false

	for i, item := range c.Items {
		if fn(item, i) {
			flag = true
			break
		}
	}

	return flag
}

func (c *OneDimensionalCollection[K, T]) DoesntContain(fn func(T, int) bool) bool {
	flag := false

	for i, item := range c.Items {
		if fn(item, i) {
			flag = true
			break
		}
	}

	return !flag
}

func (c *OneDimensionalCollection[K, T]) Diff(elements []T) Collection[K, T] {
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

	return &OneDimensionalCollection[K, T]{Items: items}
}

// Move this to maps collection
func (c *OneDimensionalCollection[K, T]) DiffAssoc(elements map[any]any) map[any]T {
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

func (c *OneDimensionalCollection[K, T]) Duplicates() []T {
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

func (c *OneDimensionalCollection[K, T]) DuplicatesBy(fieldName string) []any {
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

// I will need to revisit this
func (c *OneDimensionalCollection[K, T]) FlatMap(fn func(T, int) []K) Collection[K, K] {
	var items []K

	for i, item := range c.Items {
		items = append(items, fn(item, i)...)
	}

	//return &OneDimensionalCollection[K, K]{Items: items}
	return newCollection[K, K](items)
}

func (c *OneDimensionalCollection[K, T]) FlatMapAny(fn func(T, int) []any) []any {
	var items []any

	for i, item := range c.Items {
		items = append(items, fn(item, i)...)
	}

	return items
}

func (c *OneDimensionalCollection[K, T]) CollectFirst() T {
	var empty T

	if len(c.Items) > 0 {
		return c.Items[0]
	}

	return empty
}

func (c *OneDimensionalCollection[K, T]) First() interface{} {
	return c.CollectFirst()
}

func (c *OneDimensionalCollection[K, T]) CollectFirstWhere(fn func(T, int) bool) T {
	var item T

	for i, v := range c.Items {
		if fn(v, i) {
			item = v
			break
		}
	}

	return item
}

func (c *OneDimensionalCollection[K, T]) FirstWhere(fn func(T, int) bool) interface{} {
	return c.CollectFirstWhere(fn)
}

func (c *OneDimensionalCollection[K, T]) CollectLast() T {
	var empty T

	if len(c.Items) > 0 {
		return c.Items[len(c.Items)-1]
	}

	return empty
}

func (c *OneDimensionalCollection[K, T]) Last() interface{} {
	return c.CollectLast()
}

func (c *OneDimensionalCollection[K, T]) Count() int {
	return len(c.Items)
}

func (c *OneDimensionalCollection[K, T]) CountBy(fn func(T, int) any) map[any]int {
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

func (c *OneDimensionalCollection[K, T]) CollectAll() []T {
	return c.Items
}

func (c *OneDimensionalCollection[K, T]) All() interface{} {
	return c.Items
}

func (c *OneDimensionalCollection[K, T]) Values() []T {
	return c.Items
}

func (c *OneDimensionalCollection[K, T]) Every(fn func(T, int) bool) bool {
	flag := true

	for i, item := range c.Items {
		if !fn(item, i) {
			flag = false
			break
		}
	}

	return flag
}
