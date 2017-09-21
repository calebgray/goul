package main

import (
	"flag"
	"github.com/calebgray/goul/goul/archive"
	"github.com/calebgray/goul/goul/arguments"
)

func main() {
	arguments.Run()

	archive.Unzip(flag.Args(), "C:\\temp", arguments.Verbose)
}
