package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, min, sec int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"invalid number of time components", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", err), input}
		}
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minute: %v", err), input}
		}
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second: %v", err), input}
		}
		if hour < 0 || hour > 23 {
			return Time{}, &TimeParseError{"hour out range", fmt.Sprintf("%v", hour)}
		}
		if minute < 0 || minute > 59 {
			return Time{}, &TimeParseError{"minute out range", fmt.Sprintf("%v", minute)}
		}
		if second < 0 || second > 59 {
			return Time{}, &TimeParseError{"second out range", fmt.Sprintf("%v", second)}
		}
		return Time{hour, minute, second}, nil
	}
}

func main() {
	time, err := ParseTime("14:60:54")
	if err != nil {
		fmt.Printf("error %v time is %v", err, time)
		return
	}
	fmt.Printf("time is %v", time)
}
