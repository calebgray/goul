package arguments

import (
	"flag"
	"github.com/calebgray/goul/goul/version"
	"os"
)

var Verbose = false
var _verbose = flag.Bool("v", Verbose, "Be verbose.")

type command struct {
	help     string
	function func([]string) int
}

var _commands map[string]command

func AddCommand(cmd string, help string, function func([]string) int) {
	if _commands == nil {
		_commands = make(map[string]command)
	}

	_commands[cmd] = command{
		help:     help,
		function: function,
	}
}

func Run() int {
	flag.Parse()

	Verbose = *_verbose
	if Verbose {
		println(os.Args[0], "-", version.Get())
	}

	if len(os.Args) == 0 {
		flag.Usage()

		if _commands != nil {
			println()
			println("Available Commands:")
			for name, cmd := range _commands {
				println(" ", name, "\t", cmd.help)
			}
		}

		return 1
	}

	if _commands != nil {
		cmd := _commands[flag.Args()[0]]
		if cmd.function != nil {
			return cmd.function(flag.Args()[1:])
		}
	}

	return -1
}
