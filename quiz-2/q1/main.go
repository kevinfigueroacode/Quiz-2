// creating interfaces
// printing dates in different formats
package main

import (
	"fmt"
)

type dateFormatter interface {
	String() string
}

// create struct
type Format1 struct {
	year  int
	month int
	day   int
}

// create struct
type Format2 struct {
	day   int
	month int
	year  int
}

// create methods
func (f1 Format1) String() string {
	return fmt.Sprintf("%d/%d/%d", f1.day, f1.month, f1.year)
}

// create methods
func (f2 Format2) String() string {
	return fmt.Sprintf("%d-%d-%d", f2.year, f2.month, f2.day)
}
func main() {

	var date1 dateFormatter = Format1{
		day:   23,
		month: 02,
		year:  2023,
	}
	var date2 dateFormatter = Format2{
		year:  2023,
		month: 02,
		day:   23,
	}

	fmt.Println(date1.String())
	fmt.Println(date2.String())
}
