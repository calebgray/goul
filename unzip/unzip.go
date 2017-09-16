package main

import (
	"flag"
	"github.com/calebgray/goul/archive"
	"github.com/calebgray/goul/arguments"
)

func main() {
	arguments.Run()

	archive.Unzip(flag.Args(), "C:\\temp", arguments.Verbose)
}
