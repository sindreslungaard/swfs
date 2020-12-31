package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/sindreslungaard/swfs/internal"
)

func main() {

	err := internal.AssertDepsExist(
		"swfdump",
		"swfextract",
	)

	if err != nil {
		println(fmt.Sprint(err))
		return
	}

	input := flag.String("input", "", "The directory to read swf files from")
	output := flag.String("output", "", "The directory to output files from")
	workers := flag.Int("workers", 2, "The amount of concurrent workers")

	flag.Parse()

	if *input == "" || *output == "" {
		println("Missing required arguments, use 'extract -help' for more information")
		return
	}

	extractor := &internal.Extractor{}

	err = filepath.Walk(*input, extractor.Upload)

	if err != nil {
		println(fmt.Sprint(err))
		return
	}

	progress := make(chan string)
	done := make(chan bool)

	go extractor.Process(*workers, *output, progress, done)

	for {

		select {
		case fileName := <-progress:
			{
				println(fileName)
			}
		case <-done:
			{
				println("Done")
				return
			}
		}

	}

}
