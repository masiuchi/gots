package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/leekchan/timeutil"
)

func main() {
	rel := false
	useFormat := false
	format := "%b %d %H:%M:%S"

	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		if v == "-r" {
			rel = true
		} else if len(v) > 0 {
			useFormat = true
			format = v
		}
	}

	r := regexp.MustCompile(`\%\.([Ss])`)
	hires := r.MatchString(format)

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
		if useFormat {
			t, _ := time.Parse(format, text)
			return t.String()
		}
		panic("not implemented yet")
		// t, _ := httpdate.Str2Time(text, nil)
		// return time.Now().Second() - t.Second()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if rel == false {
			if hires {
				now := time.Now()
				microseconds := now.UnixNano() % int64(1000000000) / int64(1000)
				s := fmt.Sprintf("%06d", microseconds)
				f := r.ReplaceAllString(format, "%$1."+s)
				fmt.Print(timeutil.Strftime(&now, f))
			} else {
				t := time.Now()
				fmt.Print(timeutil.Strftime(&t, format))
			}
			fmt.Println("", scanner.Text())
		} else {
			text := r2.ReplaceAllStringFunc(scanner.Text(), repl)
			fmt.Println(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
