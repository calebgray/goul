package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"sync"
)

var (
	flagV = flag.Bool("v", false, "Be verbose.")
)

func main() {
	flag.Parse()
	if *flagV {
		fmt.Println("goul-unzip (v20170913)")
	}

	dir := "C:\\temp"
	for _, file := range flag.Args() {
		// Open Reader
		r, err := zip.OpenReader(file)
		if err != nil {
			fmt.Println("Can't open zip file:", err)
			continue
		}
		defer r.Close()

		// Extract Files in Parallel
		var wg sync.WaitGroup
		for _, f := range r.File {
			if f.Mode().IsDir() {
				continue
			}

			name := f.FileHeader.Name
			outpath := path.Join(dir, name)

			os.MkdirAll(path.Dir(outpath), 0755)

			wg.Add(1)
			go func(file *zip.File) {
				defer wg.Done()

				if *flagV {
					fmt.Println("Extracting:", name)
				}

				// Open a Zip Entry
				rc, err := file.Open()
				if err != nil {
					fmt.Println("Can't open a zip entry:", err)
					return
				}
				defer rc.Close()

				// Create Destination File
				w, err := os.Create(outpath)
				if err != nil {
					fmt.Println("Can't open a file to write:", err)
					return
				}
				defer w.Close()

				// Extract the File
				io.Copy(w, rc)
			}(f)
		}
		wg.Wait()

		// Finished
		fmt.Println("Extracted:", file)
	}
}
