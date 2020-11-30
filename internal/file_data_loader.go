package internal

import (
	"bufio"
	"strings"
)

const delimiter = ';'
const delimiterLength = 1

// FileDataLoader : file data loader.
type FileDataLoader struct {
	reader      *bufio.Reader
	headers     []string
	columnTypes []ColumnType
}

func MakeFileDataLoader(reader *bufio.Reader) FileDataLoader {
	return FileDataLoader{reader, []string{}, []ColumnType{}}
}

func (loader *FileDataLoader) Load() bool {
	for {
		line, err := loader.reader.ReadString('\n')
		if err != nil {
			break
		}
		if len(loader.headers) == 0 {
			loader.headers = strings.Split(strings.TrimSuffix(line, "\n"), ";")
		}
		if len(loader.columnTypes) == 0 {

		}
	}

	return true
}

func (loader FileDataLoader) GetHeaders() []string {
	return loader.headers
}

func (loader FileDataLoader) GetColumnTypes() []ColumnType {
	return []ColumnType{}
}
