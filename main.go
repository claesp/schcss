package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: missing arguments\n", os.Args[0])
		os.Exit(1)
	}

	var fn string
	if strings.Contains(os.Args[1], "-") {
		if os.Args[1] == "-h" {
			usage()
		}
	} else {
		fn = os.Args[1]
	}

	fmt.Fprintf(os.Stdout, "opening %s\n", fn)
}

func usage() {
	fmt.Fprintf(os.Stdout, "%s: [-h] css-file\n", os.Args[0])
	os.Exit(0)
}
