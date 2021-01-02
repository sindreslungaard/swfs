package extract

import (
	"fmt"
	"path/filepath"

	"github.com/sindreslungaard/swfs/internal"
)

// Execute runs the extract cli tool
func Execute(input string, output string, workers int) {

	err := internal.AssertDepsExist(
		"swfdump",
		"swfextract",
	)

	if err != nil {
		println(fmt.Sprint(err))
		return
	}

	if input == "" || output == "" {
		println("Missing required arguments, use 'swfs -help' for more information")
		return
	}

	extractor := &internal.Extractor{}

	err = filepath.Walk(input, extractor.Upload)

	if err != nil {
		println(fmt.Sprint(err))
		return
	}

	progress := make(chan string)
	done := make(chan bool)

	go extractor.Process(workers, output, progress, done)

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
