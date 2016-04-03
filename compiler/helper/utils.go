package helper

import (
	"strings"
)

func LispFnToGoName(method string) string {
	return strings.Replace(method, "/", ".", -1)
}
