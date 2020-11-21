package internal

import "bufio"

// DataLoader : interface for data loading.
type DataLoader interface {
	Load(reader *bufio.Reader) bool
	GetHeaders() []string
	GetColumnTypes() []ColumnType
	//get_data() [][]Type
}
