package parser

import (
	"regexp"
	"strconv"
)

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func extractFloat(contents []byte, re *regexp.Regexp) float64 {
	float, err := strconv.ParseFloat(extractString(contents, re), 64)
	if err != nil {
		return 0
	} else {
		return float
	}
}
