package main

import (
	"flag"
	"github.com/calebgray/goul/goul/archive"
	"github.com/calebgray/goul/goul/arguments"
)

func main() {
	arguments.Run()

	if flag.NArg() > 1 {
		archive.Unzip(flag.Args()[0:flag.NArg()-1], flag.Args()[flag.NArg()-1], arguments.Verbose)
	} else {
		archive.Unzip(flag.Args(), ".", arguments.Verbose)
	}
}
