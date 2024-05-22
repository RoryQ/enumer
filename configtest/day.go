// Simple test: enumeration of type int starting at 1.

package main

import "fmt"

type Day int

const (
	Monday Day = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	ck(Monday, "Monday")
	ck(Tuesday, "Tuesday")
	ck(Wednesday, "Wednesday")
	ck(Thursday, "Thursday")
	ck(Friday, "Friday")
	ck(Saturday, "Saturday")
	ck(Sunday, "Sunday")
	ck(-127, "Day(-127)")
	ck(127, "Day(127)")
}

func ck(day Day, str string) {
	if fmt.Sprint(day) != str {
		panic("day.go: " + str)
	}
}
