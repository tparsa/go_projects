package urlValidator

import (
	"regexp"
)

func Validate(url string) bool {
	match, _ := regexp.MatchString("(([a-z0-9A-Z]+)[.])?([a-z0-9A-Z]+)[.]([a-zA-Z]+)", url)
	return match
}
