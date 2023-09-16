package main

import (
	"fmt"
	str "github.com/oseintow/go-support/string"

	"github.com/oseintow/go-support/collection"
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

	c := collection.Of(employees).
		Filter(func(employee Employee, _ int) bool {
			return employee.Age > 30
		}).
		Map(func(employee Employee, _ int) Employee {
			employee.Age += 10
			return employee
		})

	fmt.Println(c.All())
	fmt.Println(c.Count())

	//avg := c.Avg(func(employee Employee) float64 {
	//	return float64(employee.Age)
	//})

	//avg := c.Av("Age")
	//fmt.Printf("avg: %v\n", avg)
	//
	//avg = c.Av(func(employee Employee) float64 {
	//	return employee.Age
	//})
	//fmt.Printf("avg 2: %v\n", avg)

	numbers := []int{2, 3, 4, 5, 6, 7, 7}

	n := collection.Of(numbers).
		Filter(func(i int, _ int) bool {
			//fmt.Println(j)
			return i <= 5
		}).
		Map(func(i int, _ int) int {
			i = i + 10
			return i
		})

	fmt.Println(n.All())

	avg := n.Avg(func(i int) float64 {
		return float64(i)
	})

	fmt.Println(avg)

	customers()
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

	r := collection.Of(customers).Chunk(4)
	for i, v := range r {
		fmt.Println(i)
		for j, s := range v {
			fmt.Printf("J is %v with value %v\n", j, s.Name)
		}
	}
	fmt.Printf("Customer avg age is %v\n", r)

	k := collection.Of([]int{2, 3, 4, 5, 6, 7, 7}).Chunk(5)
	fmt.Printf("Numbers age is %v\n", k)
}
