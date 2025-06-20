package utils

import (
	"strings"
)

func SplitString(s string, separator ...string) []string {
	sep := ","
	if len(separator) > 0 {
		sep = separator[0]
	}
	return strings.Split(s, sep)
}
