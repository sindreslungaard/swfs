package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sindreslungaard/swfs/cmd/bundle"
	"github.com/sindreslungaard/swfs/cmd/extract"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Print(`Available tools:
- swfs extract
- swfs bundle`)
		return
	}

	switch os.Args[len(os.Args)-1] {

	case "extract":
		{
			input := flag.String("input", "", "The directory to read swf files from")
			output := flag.String("output", "", "The directory to output files from")
			workers := flag.Int("workers", 2, "The amount of concurrent workers")
			flag.Parse()

			println(*input)

			extract.Execute(*input, *output, *workers)
		}

	case "bundle":
		{
			input := flag.String("input", "", "The directory with extracted swfs to convert")
			workers := flag.Int("workers", 5, "The amount of concurrent workers")
			flag.Parse()

			bundle.Execute(*input, *workers)
		}

	case "-help":
		{

			fmt.Print(`Argument list:
- swfs -input -output [-workers] extract
- swfs -input [-workers] bundle
		`)

		}

	default:
		{
			fmt.Printf("Tool %s does not exist", os.Args[1])
			return
		}

	}

}
