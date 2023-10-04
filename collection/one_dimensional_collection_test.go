package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollection_Filter(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employeesCount := Of[Employee](Employees).
			Filter(func(employee Employee, _ int) bool {
				return employee.Age > 40
			}).
			Count()

		assert.Equal(t, 2, employeesCount)
	})

	t.Run("array", func(t *testing.T) {
		count := Of[int]([]int{2, 55, 8, 3}).
			Filter(func(v int, i int) bool {
				return i > 2
			}).
			Count()

		assert.Equal(t, 1, count)
	})
}

func TestCollection_Reject(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employeesCount := Of[Employee](Employees).
			Reject(func(employee Employee, _ int) bool {
				return employee.Age > 40
			}).
			Count()

		assert.Equal(t, 4, employeesCount)
	})

	t.Run("array", func(t *testing.T) {
		count := Of[int]([]int{2, 55, 8, 3}).
			Reject(func(v int, i int) bool {
				return i > 2
			}).
			Count()

		assert.Equal(t, 3, count)
	})
}

func TestCollection_Map(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employee := Of[Employee](Employees).
			Map(func(employee Employee, _ int) Employee {
				employee.Age += 5
				return employee
			}).
			First().(Employee)

		assert.Equal(t, 89, employee.Age)
	})

	t.Run("array", func(t *testing.T) {
		val := Of[int]([]int{2, 55, 8, 3}).
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
		employees := Of[Employee](Employees).
			Chunk(2)

		assert.Equal(t, 3, len(employees))
	})

	t.Run("array", func(t *testing.T) {
		val := Of[int]([]int{2, 55, 8, 3}).
			Chunk(3)

		assert.Equal(t, 2, len(val))
	})
}

func TestCollection_First(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employee := Of[Employee](Employees).
			First().(Employee)

		assert.Equal(t, "Michael", employee.Name)
	})

	t.Run("1D struct", func(t *testing.T) {
		employee := Of1D[Employee](Employees).
			CollectFirst()

		assert.Equal(t, "Michael", employee.Name)
	})

	t.Run("array", func(t *testing.T) {
		val := Of[int]([]int{2, 55, 8, 3}).
			First()

		assert.Equal(t, 2, val)
	})
}

func TestCollection_Combine(t *testing.T) {
	slices1 := []string{"Name", "Age"}
	slices2 := []string{"Michael", "18"}

	t.Run("struct", func(t *testing.T) {
		slices := Of[string](slices1).
			Combine(slices2)

		assert.Equal(t, map[any]string{"Name": "Michael", "Age": "18"}, slices)
	})

	t.Run("array", func(t *testing.T) {
		val := Of[int]([]int{2, 10, 4, 3}).
			Combine([]int{2, 55, 8, 9})

		assert.Equal(t, map[any]int{2: 2, 10: 55, 4: 8, 3: 9}, val)
	})
}

func TestCollection_Contains(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		flag := Of[Employee](Employees).
			Contains(func(employee Employee, _ int) bool {
				return employee.Age > 400
			})

		assert.Equal(t, false, flag)
	})

	t.Run("array", func(t *testing.T) {
		flag := Of[int]([]int{2, 55, 8, 3}).
			Contains(func(v int, i int) bool {
				return v > 3
			})

		assert.Equal(t, true, flag)
	})
}

func TestCollection_DoesntContain(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		flag := Of[Employee](Employees).
			DoesntContain(func(employee Employee, _ int) bool {
				return employee.Age == 44
			})

		assert.Equal(t, true, flag)
	})

	t.Run("array", func(t *testing.T) {
		flag := Of[int]([]int{1, 2, 3, 4, 5}).
			DoesntContain(func(v int, i int) bool {
				return v < 4
			})

		assert.Equal(t, false, flag)
	})
}

func TestCollection_CountBy(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		slice := Of[Employee](Employees).
			CountBy(func(employee Employee, _ int) any {
				return employee.YearsInCompany
			})

		assert.Equal(t, map[any]int{3: 1, 6: 2, 8: 2, 10: 1}, slice)
	})

	t.Run("array", func(t *testing.T) {
		slice := Of[int]([]int{2, 55, 8, 3, 55, 3}).
			CountBy(nil)

		assert.Equal(t, map[any]int{2: 1, 3: 2, 8: 1, 55: 2}, slice)
	})
}

func TestCollection_Diff(t *testing.T) {
	slices1 := []string{"John", "Doe", "Foo", "Bar"}
	slices2 := []string{"Foo", "Mon", "Steve", "Doe"}

	t.Run("struct", func(t *testing.T) {
		slices := Of[string](slices1).
			Diff(slices2).
			All().([]string)

		assert.Equal(t, []string{"John", "Bar"}, slices)
	})

	t.Run("array", func(t *testing.T) {
		val := Of[int]([]int{2, 10, 4, 3}).
			Diff([]int{2, 3, 8, 9}).
			All().([]int)

		assert.Equal(t, []int{10, 4}, val)
	})
}

func TestOneDimensionalCollection_Duplicates(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		slices := Of[string]([]string{"a", "b", "a", "c", "b"}).
			Duplicates()

		assert.Equal(t, []string{"a", "b"}, slices)
	})
}

func TestOneDimensionalCollection_DuplicatesBy(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		slices := Of[Employee](Employees).
			DuplicatesBy("YearsInCompany")

		assert.Equal(t, []any{6, 8}, slices)
	})
}

func TestCollection_Each(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		employees := Of[Employee](Employees).
			Each(func(employee Employee, _ int) {
				employee.Age += 40
			}).
			All()

		assert.Equal(t, Employees, employees.([]Employee))
	})

	t.Run("array", func(t *testing.T) {
		arr := Of[int]([]int{2, 55, 8, 3}).
			Each(func(v int, i int) {
				i += 2
			}).
			All()

		assert.Equal(t, []int{2, 55, 8, 3}, arr.([]int))
	})
}

func TestCollection_Every(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		flag := Of[Employee](Employees).
			Every(func(employee Employee, _ int) bool {
				return employee.Age > 0
			})

		assert.Equal(t, true, flag)
	})

	t.Run("array", func(t *testing.T) {
		flag := Of[int]([]int{2, 55, 8, 3}).
			Every(func(v int, i int) bool {
				return v < 7
			})

		assert.Equal(t, false, flag)
	})
}

func TestCollection_FlatMap(t *testing.T) {
	//t.Run("struct", func(t *testing.T) {
	//	flag := Of[Employee](Employees).
	//		FlatMap(func(employee Employee, _ int) bool {
	//			return employee.Age > 0
	//		})
	//
	//	assert.Equal(t, true, flag)
	//})

	t.Run("array", func(t *testing.T) {
		arr := Of[int]([]int{2, 4, 8, 3}).
			FlatMap(func(v int, i int) []int {
				return []int{v * 2}
			})

		assert.Equal(t, []int{4, 8, 16, 6}, arr)
	})
}
