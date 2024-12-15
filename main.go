package main

import (
	"fmt"
	"io"
)

func run(proj string, out io.Writer) error{
	// check to see if project direcory is proviided
	if proj == ""{
		return fmt.Errorf("Project directory is required")
	}

	return nil
}