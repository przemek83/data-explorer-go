package internal

import (
	"bufio"
)

const delimiter = ';'
const delimiterLength = 1

// FileDataLoader : file data loader.
type FileDataLoader struct {
}

func (loader FileDataLoader) Load(reader *bufio.Reader) bool {
	// var line string
	var err error
	for {
		_, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		//fmt.Print(line)
	}

	return true
}

func (loader FileDataLoader) GetHeaders() []string {
	return []string{}
}

func (loader FileDataLoader) GetColumnTypes() []ColumnType {
	return []ColumnType{}
}
