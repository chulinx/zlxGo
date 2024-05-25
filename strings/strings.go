package strings

import "strings"

func Replaces(s string, replace map[string]string) string {
	for k, v := range replace {
		s = strings.Replace(s, k, v, -1)
	}
	return s
}
