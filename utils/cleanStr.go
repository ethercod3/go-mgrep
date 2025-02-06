package utils

import "strings"

func CleanStr(filename string) string {
	removeNewline := strings.Trim(filename, "\n")
	removeSpace := strings.TrimSpace(removeNewline)
	return removeSpace
}
