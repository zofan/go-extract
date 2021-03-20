package extract

import (
	"github.com/zofan/go-value"
	"regexp"
	"strings"
	"time"
)

var (
	timeFormats = []string{
		time.RFC3339Nano,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		time.Stamp,

		`02.01.2006 15:04:05`,

		`01/02/2006`,
		`2006 Jan 2`,
		`2006 Jan`,

		time.Kitchen,
	}

	intRe    = regexp.MustCompile(`-?\d+`)
	floatRe  = regexp.MustCompile(`-?\d+(?:,\.-?\d)?`)
	alphaRe  = regexp.MustCompile(`(?i)[a-z]+]`)
	alnumRe  = regexp.MustCompile(`(?i)[a-z\d]+]`)
	digitsRe = regexp.MustCompile(`\d+`)
	hexRe    = regexp.MustCompile(`(?i)[a-f\d]+`)
	trueRe   = regexp.MustCompile(`(?i)\b(?:true|yes|ok|accept|accepted)\b`)
	falseRe  = regexp.MustCompile(`(?i)\b(?:false|no|unknown)\b`)
)

func Alpha(v string) string {
	return alphaRe.FindString(v)
}

func Alnum(v string) string {
	return alnumRe.FindString(v)
}

func Digits(v string) string {
	return digitsRe.FindString(v)
}

func Hex(v string) string {
	return hexRe.FindString(v)
}

func Between(v, start, end string) string {
	si := strings.Index(v, start)
	if si < 0 {
		return ``
	}

	ei := strings.Index(v[si:], end)
	if ei < 0 {
		return ``
	}

	return v[si+len(start) : si+ei]
}

func Float(v string) float64 {
	return value.ToFloat(floatRe.FindString(v))
}

func Int(v string) int64 {
	return value.ToInt(intRe.FindString(v))
}

func Bool(v string) bool {
	return value.ToBool(trueRe.FindString(v))
}

func BoolSmart(v string) bool {
	return value.ToBool(trueRe.FindString(v)) || !falseRe.MatchString(v) || Int(v) > 0
}

func Time(v string) time.Time {
	for _, f := range timeFormats {
		t, err := time.Parse(f, v)
		if err == nil {
			return t
		}
	}

	return time.Time{}
}
