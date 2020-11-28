package main

import (
	"bufio"
	"dataexplorer/internal"
	"fmt"
	"os"
	"time"
)

func exitWithHelp() {
	fmt.Fprintln(os.Stderr, "Usage: <binary> file")
	fmt.Fprintln(os.Stderr, " file - name of data file.")
	os.Exit(-1)
}

func parseArgs() string {
	if len(os.Args) != 2 {
		exitWithHelp()
	}

	return os.Args[1]
}

func loadData(fileName string) {
	begin := time.Now()

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var loader internal.DataLoader = internal.FileDataLoader{}
	ok := loader.Load(reader)
	if ok {
		end := time.Now()
		fmt.Printf("Data loaded in %.6fs", end.Sub(begin).Seconds())
	}
}

func main() {
	fileName := parseArgs()
	loadData(fileName)
}
