package internal

// ColumnString - string column struct.
type ColumnString struct {
	data []string
}

// GetType - get type of column.
func (column *ColumnString) GetType() ColumnType {
	return StringColumn
}

// GetSize - Get number of elements in column.
func (column *ColumnString) GetSize() int {
	return len(column.data)
}

// Append - append value to column end.
func (column *ColumnString) Append(value string) {
	column.data = append(column.data, value)
}

// Get - get value for given index.
func (column *ColumnString) Get(index int) string {
	return column.data[index]
}
