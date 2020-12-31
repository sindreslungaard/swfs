package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Bundler struct
type Bundler struct {
	directories []string
	jobs        chan string
}

// Upload adds to the path to directories to bundle
func (b *Bundler) Upload(path string) {

	b.directories = append(b.directories, path)

	return
}

// Process starts bundling the uploaded paths
func (b *Bundler) Process(workers int, progress chan string, done chan bool) {

	b.jobs = make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		println("Started worker ", i)
		wg.Add(1)
		go b.worker(b.jobs, progress, &wg)
	}

	for _, dir := range b.directories {
		b.jobs <- dir
	}

	close(b.jobs)

	wg.Wait()

	done <- true

}

func (b *Bundler) worker(jobs chan string, progress chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	for dir := range jobs {

		files := make(map[string][]byte)

		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

			if path == dir {
				return nil
			}

			if info.IsDir() {
				return fmt.Errorf("This path includes unexpected files")
			}

			data, err := ioutil.ReadFile(path)

			if err != nil {
				fmt.Print(err)
			}

			files[info.Name()] = data

			return nil
		})

		if err != nil {
			continue
		}

		output := []byte{}

		output = append(output, []byte("version=\n1")...)

		for name, data := range files {

			output = append(output, []byte("\n\n")...)
			output = append(output, []byte(name)...)
			output = append(output, []byte("=\n")...)
			output = append(output, []byte(strings.ReplaceAll(string(data), "\n\n", ""))...)

		}

		os.RemoveAll(dir)

		err = ioutil.WriteFile(dir+".asset", output, 0644)

		if err != nil {
			fmt.Print(err)
			continue
		}

		progress <- dir

	}

}
