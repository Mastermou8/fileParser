package main

import (
	"fileparser/internal/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "./pdf"
	outdir := "./out"

	//filesList := []string{}
	println("Reading files from directory: " + root + "\n\n\n + -----------------------------")
	files, err := os.ReadDir(root) //retuns a slice of os.DirEntry, which contains the file names in the directory
	if err != nil {
		log.Fatal(err)
	}
	//parse through the files in the directory and call the ExtractFirstPage function for each file
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.EqualFold(filepath.Ext(file.Name()), ".pdf") {
			//log.Printf("Skipping non-PDF file: %s", file.Name())
			continue
		}

		println("Processing file: " + file.Name())

		//filesList = append(filesList, file.Name())
		println("Select an option: Extract first page(1) Crop file(2)")
		var x int
		_, err := fmt.Scanln(&x)
		if err != nil {
			log.Printf("Failed reading input: %v", err)
			continue
		}
		if x == 1 {

			err = utils.ExtractFirstPage(file.Name(), root, outdir)
			if err != nil {
				log.Printf("Failed extracting first page from %s: %v", file.Name(), err)
				continue
			}
		}
		if x == 2 {
			err = utils.Crop(file.Name(), root, outdir)
			if err != nil {
				log.Printf("Failed cropping %s: %v", file.Name(), err)
				continue
			}
		}
		println("Finished processing file: " + file.Name() + "\n\n + -----------------------------")
	}

}
