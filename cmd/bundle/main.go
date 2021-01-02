package bundle

import (
	"fmt"
	"io/ioutil"

	"github.com/sindreslungaard/swfs/internal"
)

// Execute runs the bundle cli tool
func Execute(input string, workers int) {

	if input == "" {
		println("Missing required arguments, use 'swfs -help' for more information")
		return
	}

	bundler := &internal.Bundler{}

	files, err := ioutil.ReadDir(input)

	if err != nil {
		fmt.Print(err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		bundler.Upload(input + "/" + file.Name())
	}

	if err != nil {
		println(fmt.Sprint(err))
		return
	}

	progress := make(chan string)
	done := make(chan bool)

	go bundler.Process(workers, progress, done)

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
