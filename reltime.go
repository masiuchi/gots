package gots

import (
	"regexp"
	"time"
)

// RelativeTime ...
type RelativeTime struct {
	regex *regexp.Regexp
	repl  func(string) string
}

// NewRelativeTime ...
func NewRelativeTime(useFormat bool, format string) *RelativeTime {
	relTime := RelativeTime{}
	relTime.regex = getRegex()
	relTime.repl = getReplaceFunc(useFormat, format)
	return &relTime
}

// GetRelativeTimeString ...
func (relTime *RelativeTime) GetRelativeTimeString(text string) string {
	return relTime.regex.ReplaceAllStringFunc(text, relTime.repl)
}

func getReplaceFunc(useFormat bool, format string) func(string) string {
	repl := func(text string) string {
		if useFormat {
			t, _ := time.Parse(format, text)
			return t.String()
		}
		panic("not implemented yet")
		// t, _ := httpdate.Str2Time(text, nil)
		// return time.Now().Second() - t.Second()
	}
	return repl
}

func getRegex() *regexp.Regexp {
	regexText := `\b(
		\d\d[-\s\/]\w\w\w	# 21 dec 17:05
			(?:\/\d\d+)?	# 21 dec/93 17:05
			[\s:]\d\d:\d\d	#       (time part of above)
			(?::\d\d)?	#       (optional seconds)
			(?:\s+[+-]\d\d\d\d)? #  (optional timezone)
		|
		\w{3}\s+\d\d\s+\d\d:\d\d:\d\d # syslog form
		|
		\d\d\d[-:]\d\d[-:]\d\dT\d\d:\d\d:\d\d.\d+ # ISO-8601
		|
		(?:\w\w\w,?\s+)?	#       (optional Day)
		\d+\s+\w\w\w\s+\d\d+\s+\d\d:\d\d:\d\d
					# 16 Jun 94 07:29:35
			(?:\s+\w\w\w|\s[+-]\d\d\d\d)?
					#	(optional timezone)
		|
		\w\w\w\s+\w\w\w\s+\d\d\s+\d\d:\d\d
					# lastlog format
	  )\b`
	regexText = regexp.MustCompile(`#.*$`).ReplaceAllString(regexText, "")
	regexText = regexp.MustCompile(`\s+`).ReplaceAllString(regexText, "")
	return regexp.MustCompile(regexText)
}
