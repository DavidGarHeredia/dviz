package main

import (
	"flag"
)

func main() {
	var problem_type string
	var file string

	flag.StringVar(&problem_type, "type", "UNDEF", "Valid options: [machines]")
	flag.StringVar(&file, "file", "UNDEF", "E.g. /home/desktop/file.json")

	flag.Parse()
}
