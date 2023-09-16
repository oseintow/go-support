package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollection_Average(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		count := Of(Employees).
			Average(func(employee Employee) float64 {
				return float64(employee.Age)
			})

		assert.Equal(t, 39.00, count)
	})

	t.Run("array", func(t *testing.T) {
		count := Of([]int{4, 9, 2, 6}).
			Avg(func(i int) float64 {
				return float64(i)
			})

		assert.Equal(t, 5.25, count)
	})
}
