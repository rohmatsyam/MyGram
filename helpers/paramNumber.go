package helpers

import (
	"errors"
	"regexp"
)

func CheckParamIsNumber(param string) error {
	sampleRegexp := regexp.MustCompile(`\d`)
	match := sampleRegexp.MatchString(param)
	if !match {
		return errors.New("not a number")
	}
	return nil
}
