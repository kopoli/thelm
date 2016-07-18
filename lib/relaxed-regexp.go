package thelm

import "regexp"

// AsRelaxedRegexp converts spaces in a string to .*
func AsRelaxedRegexp(regex string) (ret string) {
	re := regexp.MustCompile("  *")
	ret = re.ReplaceAllString(regex, ".*")
	return
}
