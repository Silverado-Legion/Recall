package utils

import "strings"

func AutoTitle(s string) string {
	if len(s) <= 30 {
		return s
	}

	result := s[:30]
	last := strings.LastIndex(result, " ")

	if last == -1 {
		return result
	}

	return result[:last]
}