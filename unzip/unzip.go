package main

import (
	"flag"
	"github.com/calebgray/goul/goul"
)

func main() {
	goul.GoulArgumentsRun()

	goul.GoulArchiveUnzip(flag.Args(), "C:\\temp", goul.Verbose)
}
