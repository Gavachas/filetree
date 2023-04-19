package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const buffsize = 1000

func main() {
	var currDir string
	flag.StringVar(&currDir, "d", `C:\GO`, "dir")
	flag.Parse()
	currDir = strings.TrimSuffix(currDir, "\n")
	currDir = strings.TrimSuffix(currDir, "\r")
	printFiles := true

	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dirTree(f, currDir, printFiles, "")

}

func dirTree(output io.Writer, currDir string, printFiles bool, str string) error {
	fileObj, err := os.Open(currDir)
	if err != nil {
		log.Fatalf("Could not open %s: %s", currDir, err.Error())
	}
	defer fileObj.Close()
	i := -1
	for i < 0 {

		files, err := fileObj.ReadDir(buffsize)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Could not read dir names in %s: %s", currDir, err.Error())
		}
		for _, file := range files {
			if file.IsDir() {
				fmt.Fprintf(output, "%s\n", str+file.Name())
				newDir := filepath.Join(currDir, file.Name())
				str += "--"
				dirTree(output, newDir, printFiles, str)
			} else if printFiles {
				fmt.Fprintf(output, "%s \n", str+file.Name())

			}

		}
	}
	return nil
}
