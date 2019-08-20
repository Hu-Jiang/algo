package easy

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Given a string date representing a Gregorian calendar date formatted
// as YYYY-MM-DD, return the day number of the year.
//
// Example 1:
// Input: date = "2019-01-09"
// Output: 9
// Explanation: Given date is the 9th day of the year in 2019.
//
// Example 2:
// Input: date = "2019-02-10"
// Output: 41
//
// Example 3:
// Input: date = "2003-03-01"
// Output: 60
//
// Example 4:
// Input: date = "2004-03-01"
// Output: 61
//
// Constraints:
// date.length == 10
// date[4] == date[7] == '-', and all other date[i]'s are digits
// date represents a calendar date between Jan 1st, 1900 and Dec 31, 2019.

func dayOfYear(date string) int {
	year, month, day, err := strToYYMMDD(date)
	if err != nil {
		log.Printf("strToYYMMDD err: %v", err)
	}

	daysOfMonth := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeapYear(year) {
		daysOfMonth[2] = 29
	}

	var res int
	for i := 1; i < month; i++ {
		res += daysOfMonth[i]
	}
	res += day

	return res
}

func strToYYMMDD(str string) (year, month, day int, err error) {
	atoi := func(str string) int {
		if err != nil {
			return 0
		}
		var res int
		res, err = strconv.Atoi(str)
		return res
	}

	dates := strings.Split(str, "-")
	if len(dates) != 3 {
		err = fmt.Errorf("%q format invalid", str)
		return
	}
	year = atoi(dates[0])
	month = atoi(dates[1])
	day = atoi(dates[2])
	if err != nil {
		err = fmt.Errorf("atoi: %v", err)
		return
	}
	return
}

func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}
