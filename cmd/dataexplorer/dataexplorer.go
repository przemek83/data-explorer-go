package main

import (
	"bufio"
	"dataexplorer/internal"
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

func parseArgs(args []string) (bool, []string) {
	if len(args) != 5 {
		return false, []string{}
	}
	return true, args[1:]
}

func createDataset(fileName string) internal.Dataset {
	begin := time.Now()

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var loader internal.DataLoader = internal.NewFileDataLoader(bufio.NewReader(file))
	dataset, err := internal.MakeDataset(loader)
	end := time.Now()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data loaded in %.6fs", end.Sub(begin).Seconds())
	return dataset
}

func main() {
	ok, params := parseArgs(os.Args)
	if !ok {
		programUsage()
		os.Exit(1)
	}
	fmt.Println("Executing program with params", params)
	createDataset(params[0])
}
