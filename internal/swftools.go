package internal

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// SWF struct
type SWF struct {
	path string
	name string
}

// Definition struct
type Definition struct {
	arg string
	ext string
}

// Extractor struct
type Extractor struct {
	files  []SWF
	jobs   chan SWF
	stdout io.Writer
}

// Upload adds to the list of files to extract
func (e *Extractor) Upload(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	if !strings.Contains(path, ".swf") {
		return nil
	}

	e.files = append(e.files, SWF{
		path: path,
		name: strings.ReplaceAll(info.Name(), ".swf", ""),
	})

	return nil
}

// Process extracts the uploaded files and sends the result to the receiver channel
// once finished, done channel is sent a message
func (e *Extractor) Process(workers int, output string, progress chan string, done chan bool) {

	e.jobs = make(chan SWF)

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		println("Started worker ", i)
		wg.Add(1)
		go e.worker(e.jobs, progress, output, &wg)
	}

	for _, file := range e.files {
		e.jobs <- file
	}

	close(e.jobs)

	wg.Wait()

	done <- true

}

func (e *Extractor) worker(jobs chan SWF, progress chan string, output string, wg *sync.WaitGroup) {

	defer wg.Done()

	for file := range jobs {

		out, err := exec.Command("swfdump", file.path).Output()

		if err != nil {
			println(fmt.Sprint(err))
			return
		}

		definitions := make(map[string]*Definition)
		exports := make(map[string]string)

		lines := strings.Split(string(out), "\n")
		for l := range lines {

			line := strings.ReplaceAll(lines[l], "\r", "")

			if strings.Contains(line, "defines id") {

				def := &Definition{}

				if strings.Contains(line, "DEFINEBINARY") {
					def.arg = "-b"
					def.ext = "bin"
				} else if strings.Contains(line, "DEFINEBITSLOSSLESS2") {
					def.arg = "-p"
					def.ext = "png"
				} else {
					continue
				}

				id := strings.Split(strings.Split(line, "defines id ")[1], " ")[0]

				definitions[id] = def

			} else if strings.Contains(line, "exports") {

				id := strings.Split(strings.Split(line, "exports ")[1], " ")[0]
				name := strings.Split(line, "\"")[1]

				exports[id] = name

			} else {
				continue
			}

		}

		if _, err := os.Stat(fmt.Sprintf("%s/%s", output, file.name)); os.IsNotExist(err) {
			os.Mkdir(fmt.Sprintf("%s/%s", output, file.name), 0755)
		}

		var wg2 sync.WaitGroup

		for id, def := range definitions {

			name, ok := exports[id]

			if !ok {
				continue
			}

			wg2.Add(1)

			go func(id2 string, def2 *Definition, waitgroup2 *sync.WaitGroup) {
				defer waitgroup2.Done()
				_, err := exec.Command(
					"swfextract",
					def2.arg,
					id2,
					"-o",
					output+"/"+file.name+"/"+name+"."+def2.ext,
					file.path,
				).Output()

				if err != nil {
					fmt.Print(err)
				}
			}(id, def, &wg2)

		}

		wg2.Wait()

		progress <- file.name

	}

}

// AssertDepsExist tests if the external binaries specified exists
func AssertDepsExist(binaries ...string) error {

	for _, bin := range binaries {
		cmd := exec.Command(bin, "--version")

		err := cmd.Run()

		if err != nil {
			return fmt.Errorf("Failed to execute '%s' binary with error '%v'", bin, err)
		}
	}

	return nil

}
