package internal

// ColumnNumeric - numeric column struct.
type ColumnNumeric struct {
	data []int
}

// GetType - get type of column.
func (column *ColumnNumeric) GetType() ColumnType {
	return NumericColumn
}

// GetSize - Get number of elements in column.
func (column *ColumnNumeric) GetSize() int {
	return len(column.data)
}

// Append - append value to column end.
func (column *ColumnNumeric) Append(value int) {
	column.data = append(column.data, value)
}

// Get - get value for given index.
func (column *ColumnNumeric) Get(index int) int {
	return column.data[index]
}
