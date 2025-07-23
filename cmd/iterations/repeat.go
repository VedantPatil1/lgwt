package interations

import "strings"

func Repeat(s string, r int) string {
	var repeated strings.Builder
	for i := 0; i < r; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
