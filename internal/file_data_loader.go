package internal

import (
	"bufio"
	"fmt"
)

const delimiter = ';'
const delimiterLenght = 1

// FileDataLoader : file data loader.
type FileDataLoader struct {
}

func (loader FileDataLoader) Load(reader *bufio.Reader) (bool, []string) {
	var line string
	var err error
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}

	return true, make([]string, 0)
}
