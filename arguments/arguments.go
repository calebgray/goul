package arguments

import (
	"flag"
	"github.com/calebgray/goul/version"
	"os"
)

var Verbose = false

var _verbose = flag.Bool("v", false, "Be verbose.")
var _commands map[string]func([]string) int

func AddCommand(cmd string, function func([]string) int) {
	if _commands == nil {
		_commands = make(map[string]func([]string) int)
	}
	_commands[cmd] = function
}

func Run() int {
	flag.Parse()

	Verbose = *_verbose
	if Verbose {
		println(os.Args[0], "-", version.Get())
	}

	if len(flag.Args()) == 0 {
		flag.Usage()
		return 1
	}

	if _commands != nil {
		cmd := _commands[flag.Args()[0]]
		if cmd != nil {
			return cmd(flag.Args()[1:])
		}
	}

	return -1
}
