package util

import (
	"regexp"
	"strings"
)

// CondenseSpaces replaces multiple spaces with a single space and trims whitespace from the start and end of the
// given string.  For example, ' a  b  c ' becomes 'a b c'.
func CondenseSpaces(s string) string {
	re := regexp.MustCompile("  +")
	return re.ReplaceAllString(strings.TrimSpace(s), " ")
}
