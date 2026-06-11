package main

import (
	"fileparser/internal/utils"
	"log"
	"os"
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
		if err := utils.ExtractFirstPage(fileName, root); err != nil {
			log.Fatal(err)
		}
	}
}
