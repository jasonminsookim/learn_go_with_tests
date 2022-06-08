package iteration

import (
	"strings"
)


func Repeat(char string, count int) (repeated string) {
	repeated = strings.Repeat(char, count)
	return
}
