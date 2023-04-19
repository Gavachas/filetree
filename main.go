package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	currDir := "с:"
	printFiles := true
	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dirTree(f, currDir, printFiles)

}

func dirTree(output io.Writer, currDir string, printFiles bool) error {
	fileObj, err := os.Open(currDir)
	if err != nil {
		log.Fatalf("Could not open %s: %s", currDir, err.Error())
	}
	defer fileObj.Close()
	files, err := fileObj.Readdir(-1)

	if err != nil {
		log.Fatalf("Could not read dir names in %s: %s", currDir, err.Error())
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Fprintf(output, "%s\n", file.Name())
			newDir := filepath.Join(currDir, file.Name())
			dirTree(output, newDir, printFiles)
		} else if printFiles {
			fmt.Fprintf(output, "%s \n", file.Name())

		}
	}
	return nil
}
