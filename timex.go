package timex

import (
	"errors"
	"time"
)

var (
	ErrorCouldNotParse = errors.New("timex: couldn't parse value according to any supplied layout")
)

var DefaultLayouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02",
	"2006-01-02 15:04:05",
}

func ParseOneInLocation(layouts []string, value string, loc *time.Location) (time.Time, error) {
	for _, e := range layouts {
		if loc == nil {
			if t, err := time.Parse(e, value); err == nil {
				return t, nil
			}
		} else {
			if t, err := time.ParseInLocation(e, value, loc); err == nil {
				return t, nil
			}
		}
	}

	return time.Time{}, ErrorCouldNotParse
}

func ParseOne(layouts []string, value string) (time.Time, error) {
	return ParseOneInLocation(layouts, value, nil)
}

func ParseDefaultsInLocation(value string, loc *time.Location) (time.Time, error) {
	return ParseOneInLocation(DefaultLayouts, value, loc)
}

func ParseDefaults(value string) (time.Time, error) {
	return ParseOne(DefaultLayouts, value)
}
