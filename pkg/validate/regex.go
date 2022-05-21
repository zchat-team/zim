package validate

import (
	"regexp"
)

const (
	phone                 = "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	numberGt0RegexString  = `^[1-9]\d*$`
	numberGte0RegexString = `^\d+$`
)

var (
	rxPhone      = regexp.MustCompile(phone)
	rxNumberGt0  = regexp.MustCompile(numberGt0RegexString)
	rxNumberGte0 = regexp.MustCompile(numberGte0RegexString)
)
