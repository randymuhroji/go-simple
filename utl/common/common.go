package common

import (
	"strings"
	"unicode"
)

func DecodeUrlArray(s string) []string {
	return strings.Split(s, ",")
}

func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func SliceStringToInterface(s []string) []interface{} {
	var result []interface{}
	for _, v := range s {
		result = append(result, v)
	}

	return result
}
