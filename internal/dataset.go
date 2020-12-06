package internal

// Dataset - struct containing all loaded data with accompanying methods.
type Dataset struct {
	headers     []string
	columnTypes []ColumnType
	data        []Column
}

// MakeDataset - create dataset object.
func MakeDataset(loader DataLoader) (bool, Dataset) {
	if !loader.Load() {
		return false, Dataset{}
	}
	return true, Dataset{loader.GetHeaders(), loader.GetColumnTypes(), loader.GetData()}
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
	return false, ""
}

// GetColumnType - get type of column for given index.
func (dataset *Dataset) GetColumnType(id int) (bool, ColumnType) {
	return false, Unknown
}

// GetData - get data as a column object for given index.
func (dataset *Dataset) GetData(id int) (bool, Column) {
	return false, dataset.data[id]
}
