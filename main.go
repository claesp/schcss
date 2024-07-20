package main

import (
	"fmt"
	"os"
	"strings"
)

func load(filename string) []byte {
	ba, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	return ba
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: missing arguments\n", os.Args[0])
		os.Exit(1)
	}

	var fn string
	if os.Args[1] == "-h" {
		usage()
	} else {
		fn = os.Args[1]
	}

	s, err := work(load(fn))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s", s)
}

func work(data []byte) (string, error) {
	s := string(data)

	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.TrimRight(s, " ")

	return s, nil
}

func usage() {
	fmt.Fprintf(os.Stdout, "%s: [-h] css-file\n", os.Args[0])
	os.Exit(0)
}
