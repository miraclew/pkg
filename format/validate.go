package format

import "regexp"

func IsMobile(number string) bool  {
	matched, _ := regexp.MatchString(`^(1[3|4|5|7|8][0-9]\d{8})$`, number)
	return matched
}
