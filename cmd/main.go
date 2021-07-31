package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"svg/cmd/conf"
	"svg/svg"
)

// inputStream defines the input stream for the program.
func inputStream(args []string) (io.ReadCloser, error) {
	if len(args) == 0 {
		// If there are no arguments use stdin.
		return os.Stdin, nil
	}
	var f io.ReadCloser
	var err error
	if f, err = os.Open(args[0]); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s is not a file", args[0])
		} else {
			return nil, fmt.Errorf("error: %w", err)
		}
	}
	return f, nil
}

func main() {

	// Setup.
	c, err := conf.Config()
	if err != nil {
		log.Fatal(err)
	}

	// Input stream.
	in, err := inputStream(c.Args())
	if err != nil {
		log.Fatal(err)
	}

	// Running mode.
	switch c.Cmd() {
	case conf.Default:
		svg.Default(in)
	case conf.Open:
		svg.Open(in)
	}
}
