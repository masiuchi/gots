package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/masiuchi/gots"
)

func main() {
	args := gots.NewArgs()
	args.ParseArgs(os.Args)

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
	r2 := regexp.MustCompile(regexText)
	repl := func(text string) string {
		if args.UseFormat {
			t, _ := time.Parse(args.Format, text)
			return t.String()
		}
		panic("not implemented yet")
		// t, _ := httpdate.Str2Time(text, nil)
		// return time.Now().Second() - t.Second()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if args.Rel == false {
			var timeString string
			now := time.Now()
			if args.IsHires() {
				timeString = gots.GetMicroTimeString(&now, args.Format)
			} else {
				timeString = gots.GetTimeString(&now, args.Format)
			}
			fmt.Println(timeString, scanner.Text())
		} else {
			text := r2.ReplaceAllStringFunc(scanner.Text(), repl)
			fmt.Println(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
