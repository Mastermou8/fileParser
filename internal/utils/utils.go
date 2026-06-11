package utils

import (
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func ExtractFirstPage(fileName string, root string) error {
	file := root + "/" + fileName //creates the path to the file

	if len(os.Args) > 1 { //if there is an argument, use it as the file name instead of the one from the directory
		fileName = os.Args[1]
	}
	//const outputDir = "out"
	if err := os.MkdirAll(root, 0o755); err != nil { //creates the output directory if it doesn't exist
		log.Fatal(err)
	}
	if err := api.ExtractPagesFile(fileName, root, []string{"1"}, nil); err != nil {
		log.Fatal(err)
	}

	log.Printf("Extracted page 1 from %s into ./%s\n", file, root)
	return nil
}
