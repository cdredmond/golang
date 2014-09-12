package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file.WriteString("This is a test of the Go file writer.\n")
	file.WriteString("This is more testing.")
}

