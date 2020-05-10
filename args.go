package gots

import "regexp"

// Args ...
type Args struct {
	Rel       bool
	UseFormat bool
	Format    string
}

// DefaultFormat ...
const DefaultFormat = "%b %d %H:%M:%S"

// NewArgs ...
func NewArgs() *Args {
	args := Args{}
	args.Format = DefaultFormat
	return &args
}

// ParseArgs ...
func (args *Args) ParseArgs(argsSlice []string) {
	for i, v := range argsSlice {
		if i == 0 {
			continue
		}
		if v == "-r" {
			args.Rel = true
		} else if len(v) > 0 {
			args.UseFormat = true
			args.Format = v
		}
	}
}

// IsHires ...
func (args *Args) IsHires() bool {
	r := regexp.MustCompile(`\%\.([Ss])`)
	return r.MatchString(args.Format)
}
