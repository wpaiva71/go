package main

import (
	"fmt"
	"io"
	"os"
)

type printFile struct{}

func main() {
	fileNames := os.Args // accepet more than 1 file to print

	// loop to iteract with all filenames passed as an argument
	for i := 1; i < len(fileNames); i++ {
		fileName := fileNames[i] // one file at the time

		f, err := os.Open(fileName) // try to open a file
		if err != nil {             // if a error occur, exit the program
			fmt.Println("Error:", err)
			os.Exit(1)
		} else { // file exists
			pf := printFile{} // new printfile struct

			io.Copy(pf, f) // use the copy function to get from file and print to console
		}
	}
}

// printFile implements the Reader interface (byteslice)
func (printFile) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
