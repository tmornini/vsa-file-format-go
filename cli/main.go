// Package main is the entry point for the application.
package main

import (
	"fmt"
	"os"

	"github.com/tmornini/vsa-file-format-go/vsafile"
)

func main() {
	file, err := vsafile.NewFileFrom(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Print(file.String())

	os.Exit(0)
}
