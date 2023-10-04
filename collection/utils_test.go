package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils_FlatMap(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		arr := FlatMap([]int{2, 4, 8, 3}, func(v int, i int) []float64 {
			return []float64{float64(v) * 2.0}
		})

		assert.Equal(t, []float64{4.0, 8.0, 16.0, 6.0}, arr)
	})

	t.Run("string", func(t *testing.T) {
		arr := FlatMap([]int{2, 4, 8, 3}, func(v int, i int) []string {
			return []string{"s"}
		})

		assert.Equal(t, []string{"s", "s", "s", "s"}, arr)
	})
}

func TestUtils_Flatten(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		arrNumber := [][]int{
			{2, 3, 5},
			{4, 5},
			{9, 2, 4, 6},
		}

		arr := Flatten(arrNumber)

		assert.Equal(t, []int{2, 3, 5, 4, 5, 9, 2, 4, 6}, arr)
	})
}
