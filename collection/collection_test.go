package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollection_Filter(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employeesCount := Of(Employees).
			Filter(func(employee Employee, _ int) bool {
				return employee.Age > 40
			}).
			Count()

		assert.Equal(t, 2, employeesCount)
	})

	t.Run("array", func(t *testing.T) {
		count := Of([]int{2, 55, 8, 3}).
			Filter(func(v int, i int) bool {
				return i > 2
			}).
			Count()

		assert.Equal(t, 1, count)
	})
}

func TestCollection_Reject(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employeesCount := Of(Employees).
			Reject(func(employee Employee, _ int) bool {
				return employee.Age > 40
			}).
			Count()

		assert.Equal(t, 4, employeesCount)
	})

	t.Run("array", func(t *testing.T) {
		count := Of([]int{2, 55, 8, 3}).
			Reject(func(v int, i int) bool {
				return i > 2
			}).
			Count()

		assert.Equal(t, 3, count)
	})
}

func TestCollection_Map(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employee := Of(Employees).
			Map(func(employee Employee, _ int) Employee {
				employee.Age += 5
				return employee
			}).
			First()

		assert.Equal(t, 89, employee.Age)
	})

	t.Run("array", func(t *testing.T) {
		val := Of([]int{2, 55, 8, 3}).
			Map(func(v int, i int) int {
				v = v * 2
				return v
			}).
			Last()

		assert.Equal(t, 6, val)
	})
}

func TestCollection_Chunk(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employees := Of(Employees).
			Chunk(2)

		assert.Equal(t, 3, len(employees))
	})

	t.Run("array", func(t *testing.T) {
		val := Of([]int{2, 55, 8, 3}).
			Chunk(3)

		assert.Equal(t, 2, len(val))
	})
}

func TestCollection_First(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employee := Of(Employees).
			First()

		assert.Equal(t, "Michael", employee.Name)
	})

	t.Run("array", func(t *testing.T) {
		val := Of([]int{2, 55, 8, 3}).
			First()

		assert.Equal(t, 2, val)
	})
}
