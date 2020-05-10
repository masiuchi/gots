package gots

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
func (a *Args) ParseArgs(args []string) {
	for i, v := range args {
		if i == 0 {
			continue
		}
		if v == "-r" {
			a.Rel = true
		} else if len(v) > 0 {
			a.UseFormat = true
			a.Format = v
		}
	}
}
