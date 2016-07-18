package thelm

import "regexp"

// AsRelaxedRegexp converts spaces in a string to .* and makes it case
// insensitive regexp
func AsRelaxedRegexp(regex string) (ret string) {
	re := regexp.MustCompile("  *")
	ret = "(?i)" + re.ReplaceAllString(regex, ".*")
	return
}
