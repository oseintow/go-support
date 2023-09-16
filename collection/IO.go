package collection

import (
	"fmt"
	"reflect"
)

type IO[V any, K any] struct {
	Values []V
	Keys   []K
}

//func OfIO[V any, K any](d []V) *IO[V, K] {
//	return NewIO(d)
//}

func NewIO[V any, K any](items []V, key []K) *IO[V, K] {
	return &IO[V, K]{
		Values: items,
	}
}

func (c *IO[V, K]) Filter(filterFunc func(u V, c any) bool) *IO[V, K] {
	var items []V

	for i, item := range c.Values {
		//r := reflect.ValueOf(i).Type()
		fmt.Println(reflect.ValueOf(i).Type())
		if filterFunc(item, i) {
			items = append(items, item)
		}
	}

	c.Values = items

	return c
}
