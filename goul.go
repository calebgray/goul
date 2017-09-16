package main

//go:generate go run goul.go build

import (
	"flag"
	"github.com/calebgray/goul/arguments"
	"os"
)

func main() {
	arguments.AddCommand("build", func(args []string) int {
		println("`command -v git > /dev/null && git rev-parse --short HEAD || echo undefined`")
		return 0
	})

	if arguments.Run() == 0 {
		os.Exit(0)
	}

	for _, arg := range flag.Args() {
		println("Unknown argument:", arg)
	}
	os.Exit(1)
}
