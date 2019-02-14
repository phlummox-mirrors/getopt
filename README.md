# getopt [![godoc](https://godoc.org/git.sr.ht/~sircmpwn/getopt?status.svg)](https://godoc.org/git.sr.ht/~sircmpwn/getopt) [![builds.sr.ht status](https://builds.sr.ht/~sircmpwn/getopt.svg)](https://builds.sr.ht/~sircmpwn/getopt)

A POSIX-compatible getopt implementation for Go, because POSIX getopt is The
Correct Way to interpret arguments to command line utilities.

Please send patches/bugs/feedback to
[~sircmpwn/public-inbox@lists.sr.ht](https://lists.sr.ht/~sircmpwn/public-inbox).

## Example Usage

```go
import (
	"os"
	"git.sr.ht/~sircmpwn/getopt"
)

func main() {
	opts, optind, err := getopt.Getopts(os.Args[1:], "abc:d:")
	if err != nil {
		panic(err)
	}
	for _, opt := range opts {
		switch opt.Option {
		case 'a':
			println("Option -a specified")
		case 'b':
			println("Option -b specified")
		case 'c':
			println("Option -c specified: " + opt.Value)
		case 'c':
			println("Option -c specified: " + opt.Value)
		}
	}
	println("Remaining arguments:")
	for _, arg := os.Args[optind:] {
		println(arg)
	}
}
```

## Future directions

At some point I intend to implement an interface similar to the
[flag](https://golang.org/pkg/flag/) package, but which interprets `os.Args`
according to POSIX getopt.
