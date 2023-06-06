package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Timepiece struct {
	hour, minute, second int
}

type ParsedError struct {
	msg, input string
}

func (p *ParsedError) Error() string {
	return fmt.Sprintf("%v: %v", p.msg, p.input)
}

func ParseTime(time string) (Timepiece, error) {
	parse := strings.Split(time, ":")

	if len(parse) != 3 {
		return Timepiece{}, &ParsedError{"Invalid number of time components", time}
	} else {
		hour, err := strconv.Atoi(parse[0])
		if err != nil {
			return Timepiece{}, &ParsedError{fmt.Sprintf("Error parsing hour: %v", err), time}
		}

		minute, err := strconv.Atoi(parse[1])
		if err != nil {
			return Timepiece{}, &ParsedError{fmt.Sprintf("Error parsing minutes: %v", err), time}
		}

		second, err := strconv.Atoi(parse[2])
		if err != nil {
			return Timepiece{}, &ParsedError{fmt.Sprintf("Error parsing second: %v", err), time}
		}

		if hour > 23 || hour < 0 {
			return Timepiece{}, &ParsedError{"Hour out of range: 0 <= hour < 23", fmt.Sprintf("%v", hour)}
		}

		if minute > 59 || minute < 0 {
			return Timepiece{}, &ParsedError{"minute out of range: 0 <= minute < 59", fmt.Sprintf("%v", minute)}
		}

		if second > 59 || second < 0 {
			return Timepiece{}, &ParsedError{"Second out of range: 0 <= second < 23", fmt.Sprintf("%v", second)}
		}

		return Timepiece{hour, minute, second}, nil
	}
}
