package validate

import (
	"github.com/shopspring/decimal"
)

// IsMobile 验证是否是手机号
func IsMobile(s string) bool {
	return rxPhone.MatchString(s)
}

func IsDecimalGt(s, t string) bool {
	d, tt, err := parseString2Decimal(s, t)
	if err != nil {
		return false
	}
	return d.GreaterThan(tt)
}

func IsDecimalGte(s, t string) bool {
	d, tt, err := parseString2Decimal(s, t)
	if err != nil {
		return false
	}
	return d.GreaterThanOrEqual(tt)
}

func IsDecimalLt(s, t string) bool {
	d, tt, err := parseString2Decimal(s, t)
	if err != nil {
		return false
	}
	return d.LessThan(tt)
}

func IsDecimalLte(s, t string) bool {
	d, tt, err := parseString2Decimal(s, t)
	if err != nil {
		return false
	}
	return d.LessThanOrEqual(tt)
}

func IsNumberGt0(s string) bool {
	return rxNumberGt0.MatchString(s)
}
func IsNumberGte0(s string) bool {
	return rxNumberGte0.MatchString(s)
}

func parseString2Decimal(s, t string) (d decimal.Decimal, tt decimal.Decimal, err error) {
	d, err = decimal.NewFromString(s)
	if err != nil {
		return
	}
	tt, err = decimal.NewFromString(t)
	if err != nil {
		panic(err.Error())
	}
	return
}
