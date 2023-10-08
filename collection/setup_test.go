package collection

import (
	"os"
	"testing"
)

type Employee struct {
	Name           string
	Age            int
	YearsInCompany int
	Sex            string
}

var Employees = []Employee{
	{Name: "Michael", Age: 84, YearsInCompany: 10, Sex: "Male"},
	{Name: "Diana", Age: 28, YearsInCompany: 3, Sex: "Female"},
	{Name: "Samuel", Age: 23, YearsInCompany: 6, Sex: "Male"},
	{Name: "Emmanuel", Age: 37, YearsInCompany: 8, Sex: "Male"},
	{Name: "Patricia", Age: 43, YearsInCompany: 8, Sex: "Female"},
	{Name: "Josephine", Age: 19, YearsInCompany: 6, Sex: "Female"},
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
