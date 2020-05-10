package gots

import (
	"fmt"
	"regexp"
	"time"

	"github.com/leekchan/timeutil"
)

// GetTimeString ...
func GetTimeString(t *time.Time, format string) string {
	return timeutil.Strftime(t, format)
}

// GetMicroTimeString ...
func GetMicroTimeString(t *time.Time, format string) string {
	microseconds := t.UnixNano() % int64(1000000000) / int64(1000)
	s := fmt.Sprintf("%06d", microseconds)
	r := regexp.MustCompile(`\%\.([Ss])`)
	f := r.ReplaceAllString(format, "%$1."+s)
	return GetTimeString(t, f)
}
