package util

import (
	"fmt"
	"math/rand"
	"time"
)

func HumanTime(secs int64) string {
	var num int64
	var quantifier string
	suffix := "后"
	diff := secs - time.Now().Unix()
	if diff < 0 {
		suffix = "前"
		diff = -diff
	}

	seconds := diff
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	months := days / 30
	years := months / 12

	switch true {
	case years > 0:
		num = years
		quantifier = "年"
	case months > 0:
		num = months
		quantifier = "月"
	case days > 0:
		num = days
		quantifier = "日"
	case hours > 0:
		num = hours
		quantifier = "时"
	case minutes > 0:
		num = minutes
		quantifier = "分"
		break
	default:
		num = seconds
		quantifier = "秒"
	}

	return fmt.Sprintf("%d%s%s", num, quantifier, suffix)
}

func RandNumber(n int) string {
	numberStr := "0123456789"
	numberBytes := []byte(numberStr)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, numberBytes[r.Intn(len(numberBytes))])
	}
	return string(result)
}

func TimeFormat(t time.Time) string {
	if t.IsZero() {
		return "0000-00-00 00:00:00"
	}
	return t.Format("2006-01-02 15:04:05")
}

func TimeToUnix(t time.Time) int64 {
	if t.IsZero() {
		return 0
	}

	return t.Unix()
}

func ShuffleString(s string) string {
	rand.Seed(time.Now().UnixNano())
	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		runes[i], runes[num] = runes[num], runes[i]
	}

	return string(runes)
}
