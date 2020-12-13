package internal

import (
	"errors"
)

// Dataset - struct containing all loaded data with accompanying methods.
type Dataset struct {
	headers     []string
	columnTypes []ColumnType
	data        []Column
}

// MakeDataset - create dataset object.
func MakeDataset(loader DataLoader) (Dataset, error) {
	if !loader.Load() {
		return Dataset{}, errors.New("Cannot load data")
	}
	return Dataset{loader.GetHeaders(), loader.GetColumnTypes(), loader.GetData()}, nil
}

// ColumnNameToID - convert column name given as string to id.
func (dataset *Dataset) ColumnNameToID(name string) (bool, int) {
	for i, currentName := range dataset.headers {
		if currentName == name {
			return true, i
		}
	}
	return false, -1
}

// ColumnIDToName - get header name for given column index.
func (dataset *Dataset) ColumnIDToName(id int) (bool, string) {
	if id >= len(dataset.headers) || id < 0 {
		return false, ""
	}
	return true, dataset.headers[id]
}

// GetColumnType - get type of column for given index.
func (dataset *Dataset) GetColumnType(id int) (bool, ColumnType) {
	if id >= len(dataset.columnTypes) || id < 0 {
		return false, Unknown
	}
	return true, dataset.columnTypes[id]
}

// GetData - get data as a column object for given index.
func (dataset *Dataset) GetData(id int) (bool, Column) {
	if id >= len(dataset.data) || id < 0 {
		return false, new(ColumnNumeric)
	}
	return true, dataset.data[id]
}
