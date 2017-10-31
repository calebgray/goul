package main

import (
	"flag"
	"github.com/calebgray/golibs/arguments"
	"os"
)

func main() {
	rc := arguments.Run()
	if rc == 0 {
		os.Exit(0)
	}

	for _, arg := range flag.Args() {
		println("Unknown argument:", arg)
	}
	os.Exit(rc)
}
