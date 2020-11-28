package internal

import (
	"bufio"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	file, err := os.Open("sample.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var loader DataLoader = FileDataLoader{}
	ok := loader.Load(reader)
	if !ok {
		t.Error("File not loaded.")
	}
}
