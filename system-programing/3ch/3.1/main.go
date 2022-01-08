package main

import (
	"io"
	"os"
	"path"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	oldFile := path.Join(pwd, "old.txt")
	file, err := os.Open(oldFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(newFile, file)
	if err != nil {
		panic(err)
	}
}
