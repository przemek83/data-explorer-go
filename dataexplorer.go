package main

import (
	"fmt"
	"os"
)

func exitWithHelp() {
	os.Exit(-1)
}

func parseArgs() string {
	if len(os.Args) != 2 {
		exitWithHelp()
	}

	return os.Args[1]
}

func main() {
	file := parseArgs()
	fmt.Println("data file: " + file)
}
