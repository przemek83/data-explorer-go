package main

import (
	"bufio"
	"dataexplorer/internal"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func programUsage() {
	fmt.Printf("Usage: %s file operation aggregation grouping\n", filepath.Base(os.Args[0]))
	fmt.Println("Where:")
	fmt.Println("\tfile        - Input file")
	fmt.Println("\toperation   - Arithmetic operation to perform")
	fmt.Println("\taggregation - Aggregation column (numerical only)")
	fmt.Println("\tgrouping    - Grouping by column")
}

func parseArgs(args []string) []string {
	if len(args) != 4 {
		programUsage()
		os.Exit(1)
	}
	return args
}

func loadData(fileName string) {
	begin := time.Now()

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	loader := internal.MakeFileDataLoader(reader)
	ok := loader.Load()
	if ok {
		end := time.Now()
		fmt.Printf("Data loaded in %.6fs", end.Sub(begin).Seconds())
	}
}

func main() {
	flag.Usage = programUsage
	params := parseArgs(os.Args[1:])
	fmt.Println("Executing program with params", params)
	loadData(params[0])
}
