package internal

import (
	"bufio"
	"fmt"
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
		line = strings.TrimSuffix(line, "\n")
		if len(loader.headers) == 0 {
			loader.headers = strings.Split(line, ";")
			continue
		}
		if len(loader.columnTypes) == 0 {
			loader.loadColumnTypes(line)
			continue
		}
	}

	return true
}

func (loader *FileDataLoader) loadColumnTypes(line string) {
	for _, columnAsString := range strings.Split(line, ";") {
		fmt.Println(columnAsString)
		loader.columnTypes = append(loader.columnTypes, ColumnTypeFromString(columnAsString))
	}
}

func (loader *FileDataLoader) GetHeaders() []string {
	return loader.headers
}

func (loader *FileDataLoader) GetColumnTypes() []ColumnType {
	return loader.columnTypes
}
