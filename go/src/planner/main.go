package main

import (
	"fmt"
	"headfirst_golang/go/src/dates"
)

func main() {
	days := 3
	fmt.Println("your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
	fmt.Println("There are 365 days a year, and", dates.DaysToWeeks(365), "weeks in a year")
	fmt.Println("100 weeks is equal to", dates.WeeksToDays(100), "days")
}
