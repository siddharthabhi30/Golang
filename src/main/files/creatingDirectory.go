package main

import (
	"fmt"
	"os"
)

func main() {
	creatingAndWritingFile()
}

func creatingDir() {
	os.Mkdir("tempFolder", os.ModePerm)
	os.Remove("tempFolder")
}

/*
we are creating, writing to a file and removing a file
*/
func creatingAndWritingFile() {
	fileName := "testing"
	ofile, _ := os.Create(fileName)
	res := "any"
	fmt.Fprintf(ofile, "%v\n", res)
	ofile.Close()
	os.Remove(fileName)
}
