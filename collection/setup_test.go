package collection

import (
	"os"
	"testing"
)

type Employee struct {
	Name           string
	Age            int
	YearsInCompany int
}

var Employees = []Employee{
	{Name: "Michael", Age: 84, YearsInCompany: 10},
	{Name: "Diana", Age: 28, YearsInCompany: 3},
	{Name: "Samuel", Age: 23, YearsInCompany: 6},
	{Name: "Emmanuel", Age: 37, YearsInCompany: 8},
	{Name: "Patricia", Age: 43, YearsInCompany: 8},
	{Name: "Josephine", Age: 19, YearsInCompany: 6},
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
