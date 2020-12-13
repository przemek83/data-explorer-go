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
	fmt.Printf("Data loaded in %.6fs\n", end.Sub(begin).Seconds())
	return dataset
}

func main() {
	dataset := createDataset(os.Args[1])
	query, err := internal.MakeQuery(os.Args[2:], &dataset)
	if err != nil {
		fmt.Println(err)
		programUsage()
		os.Exit(1)
	}
	calculator := internal.MakeCalculator(dataset)
	begin := time.Now()
	results, err := calculator.Execute(query)
	end := time.Now()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Operation completed in %.6fs\n", end.Sub(begin).Seconds())
	fmt.Printf("Results:\n%v\n", results)
}
