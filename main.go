package main

import (
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {

	root := "./pdf"
	filesList := []string{}
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filesList = append(filesList, file.Name())
	}

	for _, fileName := range filesList {
		inputFile := root + "/" + fileName

		if len(os.Args) > 1 {
			inputFile = os.Args[1]
		}

		//const outputDir = "out"
		if err := os.MkdirAll(root, 0o755); err != nil {
			log.Fatal(err)
		}

		if err := api.ExtractPagesFile(inputFile, root, []string{"1"}, nil); err != nil {
			log.Fatal(err)
		}

		log.Printf("Extracted page 1 from %s into ./%s\n", inputFile, root)
	}
}
