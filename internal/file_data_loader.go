package internal

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

const delimiter = ';'
const delimiterLength = 1

// FileDataLoader : file data loader.
type FileDataLoader struct {
	reader      *bufio.Reader
	headers     []string
	columnTypes []ColumnType
	data        []Column
}

func NewFileDataLoader(reader *bufio.Reader) *FileDataLoader {
	return &FileDataLoader{reader, []string{}, []ColumnType{}, []Column{}}
}

func (loader *FileDataLoader) Load() bool {
	for {
		line, err := loader.reader.ReadString('\n')
		if (err != nil && err != io.EOF) || len(line) == 0 {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		if len(loader.headers) == 0 {
			loader.headers = strings.Split(line, ";")
			continue
		}
		if len(loader.columnTypes) == 0 {
			loader.loadColumnTypes(line)
			if !loader.initializeData(loader.columnTypes) {
				return false
			}
			continue
		}

		if !loader.appendDataLine(line) {
			return false
		}
	}

	return loader.loadedDataOK()
}

func (loader *FileDataLoader) loadedDataOK() bool {
	for _, columnType := range loader.columnTypes {
		if columnType == Unknown {
			return false
		}
	}
	return true
}

func (loader *FileDataLoader) initializeData(columnTypes []ColumnType) bool {
	for _, columnType := range columnTypes {
		switch columnType {
		case NumericColumn:
			loader.data = append(loader.data, new(ColumnNumeric))
		case StringColumn:
			loader.data = append(loader.data, new(ColumnString))
		default:
			return false
		}
	}
	return true
}

func (loader *FileDataLoader) appendDataLine(line string) bool {
	values := strings.Split(line, ";")
	if len(values) != len(loader.data) {
		return false
	}
	for i := 0; i < len(loader.data); i++ {
		switch loader.columnTypes[i] {
		case NumericColumn:
			columnNumeric, _ := loader.data[i].(*ColumnNumeric)
			value, _ := strconv.Atoi(values[i])
			columnNumeric.Append(value)
		case StringColumn:
			columnString, _ := loader.data[i].(*ColumnString)
			columnString.Append(values[i])
		default:
			return false
		}
	}
	return true
}

func (loader *FileDataLoader) loadColumnTypes(line string) {
	for _, columnAsString := range strings.Split(line, ";") {
		loader.columnTypes = append(loader.columnTypes, ColumnTypeFromString(columnAsString))
	}
}

func (loader *FileDataLoader) GetHeaders() []string {
	return loader.headers
}

func (loader *FileDataLoader) GetColumnTypes() []ColumnType {
	return loader.columnTypes
}

func (loader *FileDataLoader) GetData() []Column {
	return loader.data
}
