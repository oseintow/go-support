package main

import (
	"fmt"
	"github.com/oseintow/go-support/collection"
	slicess "github.com/oseintow/go-support/slices"
	str "github.com/oseintow/go-support/string"
)

type Employee struct {
	Name string
	Age  float64
}

type Customer struct {
	Name string
	Age  int32
}

func main() {
	str.CanPrint()
	employees := []Employee{
		{Name: "Michael", Age: 29},
		{Name: "Dee", Age: 25},
		{Name: "Charles", Age: 45},
		{Name: "Samuel", Age: 28},
		{Name: "Diana", Age: 37},
		{Name: "Mon", Age: 33},
	}

	c := collection.Of[Employee](employees).
		Filter(func(employee Employee, _ int) bool {
			return employee.Age > 30
		}).
		Map(func(employee Employee, _ int) Employee {
			employee.Age += 10
			return employee
		}).
		All().([]Employee)

	fmt.Println(c)

	//customers()
	slices()
	sliceFlatMap()
}

func customers() {
	customers := []Customer{
		{Name: "Michael", Age: 29},
		{Name: "Dee", Age: 25},
		{Name: "Charles", Age: 45},
		{Name: "Samuel", Age: 28},
		{Name: "Diana", Age: 37},
		{Name: "Mon", Age: 33},
	}

	r := collection.Of[Customer](customers).Chunk(4)
	for i, v := range r {
		fmt.Println(i)
		for j, s := range v {
			fmt.Printf("J is %v with value %v\n", j, s.Name)
		}
	}
	fmt.Printf("Customer avg age is %v\n", r)

	k := collection.Of[int]([]int{2, 3, 4, 5, 6, 7, 7}).Chunk(5)
	fmt.Printf("Numbers age is %v\n", k)
	//t := []int{}
	//t = append(t, []int{1, 2, 3, 4}...)
	//
	//fmt.Println(t)
}

func slices() {
	//slice := slicess.Of[int]([]int{1, 2, 3, 4, 5}).
	//	Filter(func(i int, j int) bool {
	//		return i > 3
	//	}).
	//	All().([]int)
	//
	//fmt.Print("Slices: ")
	//fmt.Println(slice)

	//slice2 := slicess.Of[int]([][]int{
	//	{1, 2, 3, 4, 5},
	//	{1, 2, 3, 4, 5},
	//	{5, 4, 3, 2, 1},
	//}).
	//	Filter(func(i int, j int) bool {
	//		return i > 3
	//	}).
	//	Reject(func(i int, i2 int) bool {
	//		return i > 2
	//	}).
	//	Values().
	//	All().([][]int)
	//
	//fmt.Print("Slices: ")
	//fmt.Println(slice2)

	employees := []Employee{
		{Name: "Michael", Age: 29},
		{Name: "Dee", Age: 25},
		{Name: "Charles", Age: 45},
		{Name: "Samuel", Age: 28},
		{Name: "Diana", Age: 37},
		{Name: "Mon", Age: 33},
	}

	c := slicess.Of[Employee](employees).
		Filter(func(employee Employee, _ int) bool {
			return employee.Age > 30
		}).
		Reject(func(employee Employee, _ int) bool {
			return employee.Age < 40
		}).
		All()
	fmt.Println("Employees: ")
	fmt.Println(c)
}

func sliceFlatMap() {
	arr1 := slicess.Of[int]([]int{2, 4, 8, 3}).
		FlatMap(func(v int, i int) []int {
			return []int{2}
		})
	fmt.Println(arr1)

	arr2 := slicess.Of[float64, int]([]int{2, 4, 8, 3}).
		Filter(func(v int, _ int) bool {
			return v > 2
		}).
		//FlatMap(func(v int, i int) []int {
		//	return []int{v}
		//}).
		FlatMap(func(v int, i int) []float64 {
			return []float64{float64(v) * 2.0}
			//return []int{v}
		}).
		//Map(func(v int, _ int) int {
		//	return v * 2
		//}).
		Reject(func(v float64, _ int) bool {
			return v < 2
		}).
		All()

	fmt.Println(arr2)
}
