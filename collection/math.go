package collection

import (
	"reflect"
)

/**
'average',
'avg',
'contains',
'doesntContain',
'each',
'every',
'filter',
'first',
'flatMap',
'groupBy',
'keyBy',
'map',
'max',
'min',
'partition',
'reject',
'skipUntil',
'skipWhile',
'some',
'sortBy',
'sortByDesc',
'sum',
'takeUntil',
'takeWhile',
'unique',
'unless',
'until',
'when',
*/

func (c *OneDimensionalCollection[T]) Average(avgFunc func(T) float64) float64 {
	var total float64

	for _, v := range c.Items {
		total = total + avgFunc(v)

		// TODO: if it nil and c.Items is a linear array(ie inter or float elements) then find average
	}

	return total / float64(len(c.Items))
}

func (c *OneDimensionalCollection[T]) Avg(avgFunc func(T) float64) float64 {
	return c.Average(avgFunc)
}

func (c *OneDimensionalCollection[T]) av(t interface{}) float64 {
	var total float64

	kind := reflect.ValueOf(t).Kind()
	switch kind {
	case reflect.String:
		for _, v := range c.Items {
			y := getFieldValueFloat32(v, t.(string))
			total = total + y
		}
		return total / float64(len(c.Items))
	case reflect.Func:
		return 33333.5
	default:
		return 30.2
	}
}

func (c *OneDimensionalCollection[T]) contains() {

}
