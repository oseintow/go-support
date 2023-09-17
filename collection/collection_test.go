package collection

import (
	"fmt"
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

func TestCollection_Combine(t *testing.T) {

	slices1 := []string{"Name", "Age"}
	slices2 := []string{"Michael", "18"}

	t.Run("struct", func(t *testing.T) {
		slices := Of(slices1).
			Combine(slices2)

		fmt.Println(slices)

		assert.Equal(t, map[any]string{"Name": "Michael", "Age": "18"}, slices)
	})

	t.Run("array", func(t *testing.T) {
		val := Of([]int{2, 10, 4, 3}).
			Combine([]int{2, 55, 8, 9})

		assert.Equal(t, map[any]int{2: 2, 10: 55, 4: 8, 3: 9}, val)
	})
}

func TestCollection_Contains(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		flag := Of(Employees).
			Contains(func(employee Employee, _ int) bool {
				return employee.Age > 400
			})

		assert.Equal(t, false, flag)
	})

	t.Run("array", func(t *testing.T) {
		flag := Of([]int{2, 55, 8, 3}).
			Contains(func(v int, i int) bool {
				return v > 3
			})

		assert.Equal(t, true, flag)
	})
}

func TestCollection_CountBy(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		slice := Of(Employees).
			CountBy(func(employee Employee) any {
				return employee.YearsInCompany
			})

		assert.Equal(t, map[any]int{3: 1, 6: 2, 8: 2, 10: 1}, slice)
	})

	t.Run("array", func(t *testing.T) {
		slice := Of([]int{2, 55, 8, 3, 55, 3}).
			CountBy(nil)

		assert.Equal(t, map[any]int{2: 1, 3: 2, 8: 1, 55: 2}, slice)
	})
}
