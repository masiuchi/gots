package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/masiuchi/gots"
)

func main() {
	args := gots.NewArgs()
	args.ParseArgs(os.Args)

	// var relTime *(gots.RelativeTime)
	// if args.Rel {
	// 	relTime = gots.NewRelativeTime(args.UseFormat, args.Format)
	// }

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
			panic("-r option is not implemented yet")
			// text := relTime.GetRelativeTimeString(scanner.Text())
			// fmt.Println(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
