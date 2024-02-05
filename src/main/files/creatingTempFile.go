package main

import (
	"fmt"
	"os"
)

func main() {
	creatingTempFile()
}

func creatingTempFile() {
	oldPath := "pad.txt"
	newPath := "newfile.txt"
	tmpFile := "tempfile.txt"

	fileOrg, _ := os.Create(oldPath)
	fmt.Fprintf(fileOrg, "tempVal")
	fileOrg.Close()
	// Create a temporary file
	tempFile, err := os.CreateTemp("", tmpFile)
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}

	err = tempFile.Close()
	if err != nil {
		fmt.Println("Error closing temporary file:", err)
		return
	}

	// Copy the content of the old file to the temporary file
	oldData, err := os.ReadFile(oldPath)
	if err != nil {
		fmt.Println("Error reading old file:", err)
		return
	}
	err = os.WriteFile(tempFile.Name(), oldData, 0666)
	if err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}

	// Remove the old file
	err = os.Remove(oldPath)
	if err != nil {
		fmt.Println("Error removing old file:", err)
		return
	}

	// Rename the temporary file to the new file
	/*
	 Temp file has to be renamed else the runtime will remove it automatically.

	*/
	err = os.Rename(tempFile.Name(), newPath)
	if err != nil {
		fmt.Println("Error renaming temporary file:", err)
		return
	}
	fmt.Println("File renamed successfully")
	// comment this
	//os.Remove(newPath)
}
