package defcore

import (
	"regexp"
)

func IsValidVariableName(value string) bool {
	matched, _ := regexp.MatchString("^[a-z][a-zA-Z0-9]+$", value)
	return matched
}
