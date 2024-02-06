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

	duration, eventsPerSecond := file.EventsPerSecond()

	fmt.Print(file.String())

	fmt.Printf("Parsing duration: %s\n", duration)
	fmt.Printf("Events per second: %.2f\n", eventsPerSecond)

	os.Exit(0)
}
