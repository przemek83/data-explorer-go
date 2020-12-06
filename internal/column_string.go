package internal

// ColumnString - string column struct.
type ColumnString struct {
	data []string
}

// GetType - get type of column.
func (column *ColumnString) GetType() ColumnType {
	return StringColumn
}

// Append - append value to column end.
func (column *ColumnString) Append(value string) {
	column.data = append(column.data, value)
}

// Get - get value for given index.
func (column *ColumnString) Get(index int) string {
	return column.data[index]
}
