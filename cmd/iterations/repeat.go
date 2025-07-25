package interations

import "strings"

func Repeat(s string, r int) string {
	var repeated strings.Builder
	for range r {
		repeated.WriteString(s)
	}
	return repeated.String()
}

