package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
)

var (
	outPath = flag.String("to", ".", "Output file path")
)

func usage() {
	fmt.Print(`Usage of unzip:

  unzip [-to=out/path] file
`)
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		usage()
		return
	}

	zipFile, err := zip.OpenReader(args[0])
	check(err)
	defer zipFile.Close()

	// In the first pass, create all directories.
	for _, f := range zipFile.File {
		if f.FileInfo().IsDir() {
			folder := path.Join(*outPath, f.Name)
			os.MkdirAll(folder, f.Mode())
		}
	}

	// After creating all the directories, create all the files.
	for _, f := range zipFile.File {
		if !f.FileInfo().IsDir() {
			filePath := path.Join(*outPath, f.Name)
			if exists(filePath) {
				panic("file '" + filePath + "' already exists")
			}
			func() {
				rc, err := f.Open()
				check(err)
				defer rc.Close()

				out, err := os.Create(filePath)
				check(err)
				defer out.Close()

				_, err = io.Copy(out, rc)
				check(err)
			}()
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil || os.IsExist(err)
}
