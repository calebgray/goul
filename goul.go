package main

//go:generate go build -v ./...

import (
	"flag"
	"github.com/calebgray/goul/goul/arguments"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	arguments.AddCommand("build", func(args []string) int {
		var output []byte
		var err error
		output, err = exec.Command("git", []string{"rev-parse", "--short", "HEAD"}...).Output()
		version := "none"
		if err == nil {
			version = string(output[:7])
		}

		/*files, err := ioutil.ReadDir(".")
		if err != nil {
			return 1
		}*/
		files := []string{"archive"}
		progress := 0
		progressIncrement := 100 / len(files)
		outArg := "-o=bin/" + version + "/"
		for _, file := range files {
			progress += progressIncrement
			/*if !file.IsDir() {
				continue
			}*/
			outFile := outArg + file
			if runtime.GOOS == "windows" {
				outFile += ".exe"
			}
			exec.Command("go", []string{"build", outFile, "./" + file}...).Run()
			err = exec.Command(outFile).Run()
			if err == nil {
				print(progress, "%\n")
			} else {
				os.Remove(outFile)
			}
		}

		return 0
	})

	rc := arguments.Run()
	if rc == 0 {
		os.Exit(0)
	}

	for _, arg := range flag.Args() {
		println("Unknown argument:", arg)
	}
	os.Exit(rc)
}
