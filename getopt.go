// getopt is a POSIX-compatible implementation of getopt(3) for Go.
//
// Example usage:
//
//	import (
//		"os"
//		"git.sr.ht/~sircmpwn/getopt"
//	)
//
//	func main() {
//		opts, optind, err := getopt.Getopts(os.Args[1:], "abc:d:")
//		if err != nil {
//			panic(err)
//		}
//		for _, opt := range opts {
//			switch opt.Option {
//			case 'a':
//				println("Option -a specified")
//			case 'b':
//				println("Option -b specified")
//			case 'c':
//				println("Option -c specified: " + opt.Value)
//			case 'd':
//				println("Option -d specified: " + opt.Value)
//			}
//		}
//		println("Remaining arguments:")
//		for _, arg := range os.Args[optind:] {
//			println(arg)
//		}
//	}
package getopt

import (
	"fmt"
	"os"
)

// In the case of "-o example", Option is 'o' and "example" is Value. For
// options which do not take an argument, Value is "".
type Option struct {
	Option rune
	Value  string
}

// This is returned when an unknown option is found in argv, but not in the
// option spec.
type UnknownOptionError rune

func (e UnknownOptionError) Error() string {
	return fmt.Sprintf("%s: unknown option -%c", os.Args[0], rune(e))
}

// This is returned when an option with a mandatory argument is missing that
// argument.
type MissingOptionError rune

func (e MissingOptionError) Error() string {
	return fmt.Sprintf("%s: expected argument for -%c", os.Args[0], rune(e))
}

// Getopts implements a POSIX-compatible options interface.
//
// Returns a slice of options and the index of the first non-option argument.
//
// If an error is returned, you must print it to stderr to be POSIX complaint.
func Getopts(argv []string, spec string) ([]Option, int, error) {
	optmap := make(map[rune]bool)
	runes := []rune(spec)
	for i, rn := range spec {
		if rn == ':' {
			if i == 0 {
				continue
			}
			optmap[runes[i-1]] = true
		} else {
			optmap[rn] = false
		}
	}

	var (
		i    int
		opts []Option
	)
	for i = 1; i < len(argv); i++ {
		arg := argv[i]
		runes = []rune(arg)
		if len(arg) == 0 || arg == "-" {
			break
		}
		if arg[0] != '-' {
			break
		}
		if arg == "--" {
			i++
			break
		}
		for j, opt := range runes[1:] {
			if optopt, ok := optmap[opt]; !ok {
				opts = append(opts, Option{'?', ""})
				return opts, i, UnknownOptionError(opt)
			} else if optopt {
				if j+1 < len(runes)-1 {
					opts = append(opts, Option{opt, string(runes[j+2:])})
					break
				} else {
					if i+1 >= len(argv) {
						if len(spec) >= 1 && spec[0] == ':' {
							opts = append(opts, Option{':', string(opt)})
						} else {
							return opts, i, MissingOptionError(opt)
						}
					} else {
						opts = append(opts, Option{opt, argv[i+1]})
						i++
					}
				}
			} else {
				opts = append(opts, Option{opt, ""})
			}
		}
	}
	return opts, i, nil
}
