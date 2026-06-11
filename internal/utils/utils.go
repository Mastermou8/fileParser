package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ExtractFirstPage(fileName string, root string, outdir string) error {
	println("Extracting first page from file: " + fileName + "\n")
	file := root + "/" + fileName //creates the path to the file

	if len(os.Args) > 1 { //if there is an argument, use it as the file name instead of the one from the directory
		fileName = os.Args[1]
	}

	if err := os.MkdirAll(outdir, 0o755); err != nil { //creates the output directory if it doesn't exist
		return fmt.Errorf("create output dir: %w", err)
	}
	if err := api.ExtractPagesFile(file, outdir, []string{"1"}, nil); err != nil {
		return fmt.Errorf("extract first page: %w", err)
	}

	//log.Printf("Extracted page 1 from %s into ./%s\n", file, root)

	return nil
}

func Crop(fileName string, root string, outdir string) error {
	box := &model.Box{
		Rect: types.RectForWidthAndHeight(0, 400, 600, 400),
	}
	println("cropping first page from file: " + fileName + "\n")
	file := root + "/" + fileName //creates the path to the file

	if len(os.Args) > 1 { //if there is an argument, use it as the file name instead of the one from the directory
		fileName = os.Args[1]
	}

	if err := os.MkdirAll(outdir, 0o755); err != nil { //creates the output directory if it doesn't exist
		return fmt.Errorf("create output dir: %w", err)
	}

	cropOutFile := filepath.Join(outdir, "cropped_"+fileName)
	if err := api.CropFile(file, cropOutFile, []string{"1"}, box, nil); err != nil {
		return fmt.Errorf("crop file: %w", err)
	}

	//log.Printf("Cropped page 1 from %s into %s\n", file, cropOutFile)

	return nil
}
