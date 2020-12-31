package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/sindreslungaard/swfs/internal"
)

func main() {

	input := flag.String("input", "", "The directory with extracted swfs to convert")
	workers := flag.Int("workers", 5, "The amount of concurrent workers")

	flag.Parse()

	if *input == "" {
		println("Missing required arguments, use 'bundle -help' for more information")
		return
	}

	bundler := &internal.Bundler{}

	files, err := ioutil.ReadDir(*input)

	if err != nil {
		fmt.Print(err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		bundler.Upload(*input + "/" + file.Name())
	}

	if err != nil {
		println(fmt.Sprint(err))
		return
	}

	progress := make(chan string)
	done := make(chan bool)

	go bundler.Process(*workers, progress, done)

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
