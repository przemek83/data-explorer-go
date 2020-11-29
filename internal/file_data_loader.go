package internal

import (
	"bufio"
)

const delimiter = ';'
const delimiterLength = 1

// FileDataLoader : file data loader.
type FileDataLoader struct {
	reader *bufio.Reader
}

func MakeFileDataLoader(reader *bufio.Reader) FileDataLoader {
	return FileDataLoader{reader}
}

func (loader FileDataLoader) Load() bool {
	// var line string
	var err error
	for {
		_, err = loader.reader.ReadString('\n')
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
