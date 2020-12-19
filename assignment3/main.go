package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Check to see that there is an additional argument to the command line
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	// Get the filename from the command line arguments
	fName := os.Args[1]

	// Open the file with the os.Open function
	f, err := os.Open(fName)

	// Check to see there was no error opening the file
	if err != nil {
		fmt.Println("There was an error:", err)
		os.Exit(1)
	}

	// Print the file contents to the stdout
	io.Copy(os.Stdout, f)
}
